package order

import (
	"context"
	"errors"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/shared"
)

type PlaceOrderUseCase struct {
	orderRepo  Repo
	raffleRepo raffle.Repo
	pixPayment PixPayment
	uuid       shared.UUIDGenerator
}

var (
	errReachedLimit = errors.New("user reached the limit")
	errUnavaliable  = errors.New("item unavailable")
)

func NewPlaceOrderUseCase(orderRepo Repo, raffleRepo raffle.Repo, pixPayment PixPayment, uuid shared.UUIDGenerator) *PlaceOrderUseCase {
	return &PlaceOrderUseCase{orderRepo: orderRepo, raffleRepo: raffleRepo, pixPayment: pixPayment, uuid: uuid}
}

func (u *PlaceOrderUseCase) Run(ctx context.Context, model *Request) (Order, error) {
	order := Order{
		ID:            u.uuid.Generate(),
		ProductID:     model.ProductID,
		Items:         model.Items,
		UserID:        "GUID",
		PaymentMethod: PIX,
	}

	raffleItem, err := u.raffleRepo.GetProduct(ctx, order.ProductID)
	if err != nil {
		return order, err
	}

	if raffleItem.UserLimit != 0 {
		existingOrder, errGet := u.orderRepo.GetUserOrders(ctx, order.ProductID)
		if err != nil {
			return order, errGet
		}

		if existingOrder != nil {
			return order, errReachedLimit
		}
	}

	if err != nil {
		return order, err
	}

	for _, v := range order.Items {
		if !checkAvaliability(raffleItem.Variation, v) {
			return order, errUnavaliable
		}
	}

	order.Pix, err = u.pixPayment.GeneratePix()
	if err != nil {
		return order, err
	}

	order.Total = len(order.Items) * raffleItem.UnitPrice

	_, err = u.orderRepo.CreateOrder(ctx, model)
	if err != nil {
		return order, err
	}

	return order, nil
}

func checkAvaliability(s []raffle.Variation, e int) bool {
	for _, a := range s {
		if a.Number == e && a.Status != raffle.Available {
			return true
		}
	}

	return false
}
