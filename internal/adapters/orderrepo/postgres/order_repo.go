package postgres

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/order"
	postgres "github.com/evmartinelli/go-rifa-microservice/pkg/postgres"
)

type OrderRepo struct {
	db *postgres.Database
}

func NewOrderRepo(db *postgres.Database) *OrderRepo {
	return &OrderRepo{db}
}

func (r *OrderRepo) CreateOrder(ctx context.Context, rm *order.Order) error {
	insertedOrder := &Order{
		RaffleSlug:    rm.ProductID,
		UserID:        rm.UserID,
		Total:         rm.Total,
		PaymentMethod: int(rm.PaymentMethod),
		Status:        rm.Status,
	}
	tx := r.db.Begin()
	tx.Create(&insertedOrder)

	for _, v := range rm.Items {
		tx.Create(&OrderItems{
			OrderID: insertedOrder.ID,
			Number:  v,
		})
	}

	tx.Create(&PixPayment{
		OrderID:      insertedOrder.ID,
		QRCode:       rm.Pix.QRCode,
		QRCodeBase64: rm.Pix.QRCodeBase64,
		Status:       rm.Pix.Status,
	})

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r *OrderRepo) GetUserOrders(ctx context.Context, pid string) ([]order.Order, error) {
	return nil, nil
}
