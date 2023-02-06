package dynamodb

import (
	"context"
	"fmt"
	"strconv"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/order"
	db "github.com/evmartinelli/go-rifa-microservice/pkg/dynamodb"
)

const (
	OrderPK  = "P#%v#U#%v"
	OrderSK  = "O#%v"
	ItemType = "Order"
)

type OrderRepo struct {
	db *db.DynamoConfig
}

func NewOrderRepo(mdb *db.DynamoConfig) *OrderRepo {
	return &OrderRepo{mdb}
}

func (r *OrderRepo) CreateOrder(ctx context.Context, rm *order.Order) error {
	_, err := r.db.Save(OrderToDynamo(rm))
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepo) GetUserOrders(ctx context.Context, pid string) ([]order.Order, error) {
	results := []order.Order{}

	err := r.db.Query(fmt.Sprintf(OrderPK, pid, pid), &results)
	if err != nil {
		return []order.Order{}, err
	}

	return results, nil
}

func OrderToDynamo(o *order.Order) DynamoOrder {
	return DynamoOrder{
		PK:            fmt.Sprintf(OrderPK, o.ProductID, o.UserID),
		SK:            fmt.Sprintf(OrderSK, o.ID),
		GSI1PK:        strconv.Itoa(o.Pix.ID),
		ID:            o.ID,
		ProductID:     o.ProductID,
		UserID:        o.UserID,
		Total:         o.Total,
		PaymentMethod: int(o.PaymentMethod),
		Items:         o.Items,
		Pix:           o.Pix,
		ItemType:      ItemType,
	}
}

func DynamoToOrder(dyn *DynamoOrder) order.Order {
	return order.Order{
		ID:            dyn.ID,
		ProductID:     dyn.ProductID,
		UserID:        dyn.UserID,
		Total:         dyn.Total,
		PaymentMethod: order.PaymentMethod(dyn.PaymentMethod),
		Items:         dyn.Items,
		Pix:           dyn.Pix,
	}
}
