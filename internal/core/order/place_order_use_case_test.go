package order_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/idgenerator"
	mock_order "github.com/evmartinelli/go-rifa-microservice/internal/core/mock/order"
	mock_raffle "github.com/evmartinelli/go-rifa-microservice/internal/core/mock/raffle"
	mock_shared "github.com/evmartinelli/go-rifa-microservice/internal/core/mock/shared"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/order"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/shared"
)

func placeOrderUseCase(t *testing.T) (uc *order.PlaceOrderUseCase,
	or *mock_order.MockRepo,
	rr *mock_raffle.MockRepo,
	p *mock_order.MockPixPayment,
	ur *mock_shared.MockAuth,
) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	orderRepo := mock_order.NewMockRepo(mockCtl)
	raffleRepo := mock_raffle.NewMockRepo(mockCtl)
	uuid := idgenerator.NewUUIDGenerator()
	pix := mock_order.NewMockPixPayment(mockCtl)
	userrepo := mock_shared.NewMockAuth(mockCtl)

	placeOrderUseCase := order.NewPlaceOrderUseCase(orderRepo, raffleRepo, pix, uuid, userrepo)

	return placeOrderUseCase, orderRepo, raffleRepo, pix, userrepo
}

func TestCreateOrder(t *testing.T) {
	t.Parallel()

	orderUseCase, repo, raffleRepo, pix, userrepo := placeOrderUseCase(t)

	repo.EXPECT().CreateOrder(context.Background(), gomock.Any()).AnyTimes().Return(nil)
	pix.EXPECT().GeneratePix().AnyTimes().Return(order.Pix{}, nil)
	userrepo.EXPECT().Claims(gomock.Any()).AnyTimes().Return(&shared.User{}, nil)

	t.Run("Given a product with user Limit, it returns error since the user has order", func(t *testing.T) {
		t.Parallel()

		repo.EXPECT().GetUserOrders(gomock.Any(), gomock.Any()).Return([]order.Order{{ID: "ID2"}, {ID: "ID2"}}, nil)
		raffleRepo.EXPECT().GetProduct(context.Background(), gomock.Any()).Return(raffle.Raffle{UserLimit: 1}, nil)

		expectederr := order.ErrReachedLimit

		res, err := orderUseCase.Run(context.Background(), &order.Request{}, "")
		require.Error(t, expectederr, err)
		require.Nil(t, res)
	})

	t.Run("Given a product with user Limit, it returns error since the user is buying more then allowed", func(t *testing.T) {
		t.Parallel()

		repo.EXPECT().GetUserOrders(gomock.Any(), gomock.Any()).Return([]order.Order{}, nil)
		raffleRepo.EXPECT().GetProduct(context.Background(), gomock.Any()).Return(raffle.Raffle{UserLimit: 1}, nil)
		raffleRepo.EXPECT().UpdateItems(gomock.Any(), gomock.Any()).Return(nil)

		expectederr := order.ErrReachedLimit

		res, err := orderUseCase.Run(context.Background(), &order.Request{Items: []int{1, 2}}, "")
		require.Error(t, expectederr, err)
		require.Nil(t, res)
	})

	t.Run("Given a product without user Limit, it returns a order", func(t *testing.T) {
		t.Parallel()

		repo.EXPECT().GetUserOrders(gomock.Any(), gomock.Any()).Return([]order.Order{{ID: "ID2"}, {ID: "ID2"}}, nil)
		raffleRepo.EXPECT().GetProduct(context.Background(), gomock.Any()).Return(raffle.Raffle{}, nil)
		raffleRepo.EXPECT().UpdateItems(gomock.Any(), gomock.Any()).Return(nil)

		res, err := orderUseCase.Run(context.Background(), &order.Request{ProductID: "mockID", Items: []int{1}}, "")
		require.Nil(t, err)
		require.Equal(t, int(res.PaymentMethod), int(order.PIX))
	})
}
