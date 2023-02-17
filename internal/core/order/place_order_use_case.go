package order

import (
	"context"

	auth "github.com/evmartinelli/go-rifa-microservice/internal/adapters/shared/token"
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

func (u *PlaceOrderUseCase) Run(ctx context.Context, model *Request, token string) (*Order, error) {
	claims, err := auth.Claims(token)
	if err != nil {
		return nil, err
	}

	order := &Order{
		ID:            u.uuid.Generate(),
		ProductID:     model.ProductID,
		Items:         model.Items,
		PaymentMethod: PIX,
		UserID:        claims.Username,
		Status:        string(Created),
	}

	if errOrder := u.validateOrder(ctx, order); err != nil {
		return nil, errOrder
	}

	order.Pix, err = u.pixPayment.GeneratePix()
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

func (u *PlaceOrderUseCase) validateOrder(ctx context.Context, order *Order) error {
	raffleItem, err := u.raffleRepo.GetProduct(ctx, order.ProductID)
	if err != nil {
		return err
	}

	if u.hasUserLimit(raffleItem.UserLimit) {
		if u.hasOrder(ctx, order, raffleItem.UserLimit) {
			return ErrReachedLimit
		}
	}

	for _, v := range order.Items {
		if !checkAvaliability(raffleItem.Numbers, v) {
			return ErrUnavaliable
		}
	}

	for _, v := range order.Items {
		if err := u.updateItemStatus(ctx, &raffleItem, v); err != nil {
			return err
		}
	}

	order.Total = float32(len(order.Items)) * raffleItem.UnitPrice

	if err := u.orderRepo.CreateOrder(ctx, order); err != nil {
		return err
	}

	return nil
}
