package order_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	mock_order "github.com/evmartinelli/go-rifa-microservice/internal/core/mock/order"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/order"
)

func ListOrderUseCase(t *testing.T) (uc *order.ListOrderUseCase, or *mock_order.MockRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	orderRepo := mock_order.NewMockRepo(mockCtl)

	listOrderUseCase := order.NewListOrderUseCase(orderRepo)

	return listOrderUseCase, orderRepo
}

func TestListOrder(t *testing.T) {
	t.Parallel()

	orderUseCase, repo := ListOrderUseCase(t)

	repo.EXPECT().GetUserOrders(gomock.Any(), gomock.Any()).Return([]order.Order{{ID: "ID2"}, {ID: "ID2"}}, nil)

	res, err := orderUseCase.Run(context.Background(), "123")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(res) != 2 {
		t.Fatalf("expected 2 items, got %v", len(res))
	}
}
