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
		UserID:        "F9168C5E-CEB2-4faa-B6BF-329BF39FA1E4",
		PaymentMethod: PIX,
	}

	raffleItem, err := u.raffleRepo.GetProduct(ctx, order.ProductID)
	if err != nil {
		return order, err
	}

	if u.hasUserLimit(raffleItem.UserLimit) && u.hasOrder(ctx, &order) {
		return order, errReachedLimit
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

	for _, v := range order.Items {
		err = u.updateItemStatus(ctx, &raffleItem, v)
		if err != nil {
			return order, err
		}
	}

	order.Total = len(order.Items) * raffleItem.UnitPrice

	_, err = u.orderRepo.CreateOrder(ctx, &order)
	if err != nil {
		return order, err
	}

	return order, nil
}

func checkAvaliability(s []raffle.Variation, e int) bool {
	for _, a := range s {
		if a.Number == e && a.Status != raffle.Available {
			return false
		}
	}

	return true
}

func (u *PlaceOrderUseCase) updateItemStatus(ctx context.Context, s *raffle.Raffle, v int) error {
	updateVariation := make([]raffle.Variation, 0)

	for i := range s.Variation {
		if s.Variation[i].Number == v {
			s.Variation[i].Status = raffle.Pending
			updateVariation = append(updateVariation, s.Variation[i])
		}
	}

	err := u.raffleRepo.UpdateItems(ctx, updateVariation)
	if err != nil {
		return err
	}

	return nil
}

func (u *PlaceOrderUseCase) hasUserLimit(userLimit int) bool {
	return userLimit != 0
}

func (u *PlaceOrderUseCase) hasOrder(ctx context.Context, order *Order) bool {
	existingOrder, errGet := u.orderRepo.GetUserOrders(ctx, order.ProductID)
	if errGet != nil {
		return false
	}

	if existingOrder != nil {
		return true
	}

	return false
}
