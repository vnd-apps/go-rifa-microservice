package order

import "github.com/evmartinelli/go-rifa-microservice/internal/core/shared"

type PlaceOrderUseCase struct {
	orderRepo  Repo
	pixPayment PixPayment
	uuid       shared.UUIDGenerator
}

func NewPlaceOrderUseCase(orderRepo Repo, pixPayment PixPayment, uuid shared.UUIDGenerator) *PlaceOrderUseCase {
	return &PlaceOrderUseCase{orderRepo: orderRepo, pixPayment: pixPayment, uuid: uuid}
}

func (u *PlaceOrderUseCase) Run(model *PlaceOrderRequest) (PlaceOrderResponse, error) {
	return PlaceOrderResponse{}, nil
}
