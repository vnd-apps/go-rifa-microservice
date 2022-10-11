package order_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/idgenerator"
	mock_order "github.com/evmartinelli/go-rifa-microservice/internal/core/mock/order"
	mock_raffle "github.com/evmartinelli/go-rifa-microservice/internal/core/mock/raffle"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/order"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	"github.com/evmartinelli/go-rifa-microservice/pkg/assert"
)

type test struct {
	name string
	mock func()
	err  error
}

var errReachedLimit = errors.New("user reached the limit")

func placeOrderUseCase(t *testing.T) (uc *order.PlaceOrderUseCase, or *mock_order.MockRepo, rr *mock_raffle.MockRepo, p *mock_order.MockPixPayment) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	orderRepo := mock_order.NewMockRepo(mockCtl)
	raffleRepo := mock_raffle.NewMockRepo(mockCtl)
	uuid := idgenerator.NewUUIDGenerator()
	pix := mock_order.NewMockPixPayment(mockCtl)

	placeOrderUseCase := order.NewPlaceOrderUseCase(orderRepo, raffleRepo, pix, uuid)

	return placeOrderUseCase, orderRepo, raffleRepo, pix
}

func TestCreateOrder(t *testing.T) {
	t.Parallel()

	orderUseCase, repo, raffleRepo, pix := placeOrderUseCase(t)

	tests := []test{
		{
			name: "User Reached Limit",
			mock: func() {
				repo.EXPECT().CreateOrder(context.Background(), order.Order{}).Return(order.Order{}, nil)
				repo.EXPECT().GetUserOrders(gomock.Any(), gomock.Any()).Return([]order.Order{}, nil)
				raffleRepo.EXPECT().GetProduct(context.Background(), gomock.Any()).Return(raffle.Raffle{UserLimit: 1}, nil)
				pix.EXPECT().GeneratePix().Return(order.Pix{}, nil)
			},
			err: errReachedLimit,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()
			res, err := orderUseCase.Run(context.Background(), &order.Request{})
			assert.NotNil(t, err)
			assert.NotNil(t, res)
		})
	}
}
