package order

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/shared"
)

type PlaceOrderUseCase struct {
	orderRepo  Repo
	raffleRepo raffle.Repo
	pixPayment PixPayment
	uuid       shared.UUIDGenerator
}

func NewPlaceOrderUseCase(orderRepo Repo, raffleRepo raffle.Repo, pixPayment PixPayment, uuid shared.UUIDGenerator) *PlaceOrderUseCase {
	return &PlaceOrderUseCase{orderRepo: orderRepo, raffleRepo: raffleRepo, pixPayment: pixPayment, uuid: uuid}
}

func (u *PlaceOrderUseCase) Run(ctx context.Context, model *Request, userID string) (*Order, error) {
	order := &Order{
		ID:            u.uuid.Generate(),
		ProductID:     model.ProductID,
		Items:         model.Items,
		PaymentMethod: PIX,
		UserID:        userID,
		Status:        string(Created),
	}

	raffleItem, err := u.raffleRepo.GetProduct(ctx, order.ProductID)
	if err != nil {
		return nil, err
	}

	if u.hasUserLimit(raffleItem.UserLimit) {
		if u.hasOrder(ctx, order, raffleItem.UserLimit) {
			return nil, ErrReachedLimit
		}
	}

	for _, v := range order.Items {
		if !checkAvaliability(raffleItem.Numbers, v) {
			return nil, ErrUnavaliable
		}
	}

	order.Pix, err = u.pixPayment.GeneratePix()
	if err != nil {
		return nil, err
	}

	for _, v := range order.Items {
		err = u.updateItemStatus(ctx, &raffleItem, v)
		if err != nil {
			return nil, err
		}
	}

	order.Total = float32(len(order.Items)) * raffleItem.UnitPrice

	err = u.orderRepo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func checkAvaliability(s []raffle.Numbers, e int) bool {
	for _, a := range s {
		if a.Number == e && a.Status != raffle.Available {
			return false
		}
	}

	return true
}

func (u *PlaceOrderUseCase) updateItemStatus(ctx context.Context, s *raffle.Raffle, v int) error {
	updateVariation := make([]raffle.Numbers, 0)

	for i := range s.Numbers {
		if s.Numbers[i].Number == v {
			s.Numbers[i].Status = raffle.Pending
			updateVariation = append(updateVariation, s.Numbers[i])
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

func (u *PlaceOrderUseCase) hasOrder(ctx context.Context, order *Order, userLimit int) bool {
	if len(order.Items) > userLimit {
		return true
	}

	existingOrder, errGet := u.orderRepo.GetUserOrders(ctx, order.ProductID)
	if errGet != nil {
		return false
	}

	if existingOrder != nil {
		return true
	}

	return false
}
