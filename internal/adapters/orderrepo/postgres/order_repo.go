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

	err := tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepo) GetUserOrders(ctx context.Context, userID string) ([]order.Order, error) {
	var orders []Order

	r.db.Where("user_id = ?", userID).Find(&orders)

	var ordersResponse []order.Order

	for i := range orders {
		var items []OrderItems

		r.db.Where("order_id = ?", orders[i].ID).Find(&items)

		var pix *PixPayment

		r.db.Where("order_id = ?", orders[i].ID).Find(&pix)

		ordersResponse = append(ordersResponse, order.Order{
			ProductID: orders[i].RaffleSlug,
			UserID:    orders[i].UserID,
			Total:     orders[i].Total,
			Items:     orderNumbers(items),
			Pix:       orderPix(pix),
		})
	}

	return ordersResponse, nil
}

func orderPix(pix *PixPayment) order.Pix {
	return order.Pix{
		QRCode:       pix.QRCode,
		QRCodeBase64: pix.QRCodeBase64,
		Status:       pix.Status,
	}
}

func orderNumbers(items []OrderItems) []int {
	var numbers []int
	for _, v := range items {
		numbers = append(numbers, v.Number)
	}

	return numbers
}
