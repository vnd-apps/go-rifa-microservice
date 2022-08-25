package order

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/shared"
)

type PlaceOrderUseCase struct {
	orderRepo  Repo
	pixPayment PixPayment
	uuid       shared.UUIDGenerator
}

func NewPlaceOrderUseCase(orderRepo Repo, pixPayment PixPayment, uuid shared.UUIDGenerator) *PlaceOrderUseCase {
	return &PlaceOrderUseCase{orderRepo: orderRepo, pixPayment: pixPayment, uuid: uuid}
}

func (u *PlaceOrderUseCase) Run(ctx context.Context, model *PlaceOrderRequest) (PlaceOrderResponse, error) {
	_, err := u.orderRepo.CreateOrder(ctx, model)
	if err != nil {
		return PlaceOrderResponse{}, err
	}

	return PlaceOrderResponse{}, nil
}
