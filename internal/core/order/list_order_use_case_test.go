package order_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	mock_order "github.com/evmartinelli/go-rifa-microservice/internal/core/mock/order"
	mock_shared "github.com/evmartinelli/go-rifa-microservice/internal/core/mock/shared"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/order"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/shared"
)

func ListOrderUseCase(t *testing.T) (uc *order.ListOrderUseCase, or *mock_order.MockRepo, sr *mock_shared.MockAuth) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	orderRepo := mock_order.NewMockRepo(mockCtl)
	sharedRepo := mock_shared.NewMockAuth(mockCtl)

	listOrderUseCase := order.NewListOrderUseCase(orderRepo, sharedRepo)

	return listOrderUseCase, orderRepo, sharedRepo
}

func TestListOrder(t *testing.T) {
	t.Parallel()

	orderUseCase, repo, sharedRepo := ListOrderUseCase(t)

	repo.EXPECT().GetUserOrders(gomock.Any(), gomock.Any()).Return([]order.Order{{ID: "ID2"}, {ID: "ID2"}}, nil)
	sharedRepo.EXPECT().Claims(gomock.Any()).Return(&shared.User{Username: "Evandro"}, nil)

	res, err := orderUseCase.Run(context.Background(), "123")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(res) != 2 {
		t.Fatalf("expected 2 items, got %v", len(res))
	}
}
