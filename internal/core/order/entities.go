package order

import "errors"

type Request struct {
	ProductID string `json:"productID" binding:"required" example:"30dd879c-ee2f-11db-8314-0800200c9a66"`
	Items     []int  `json:"numbers" binding:"required" example:"1,2"`
}

type Order struct {
	ID            string
	ProductID     string
	UserID        string
	Total         float32
	PaymentMethod PaymentMethod
	Items         []int
	Pix           Pix
	Status        string
}

type Pix struct {
	ID           int
	Status       string
	QRCode       string
	QRCodeBase64 string
}

type (
	PaymentMethod int
	Status        string
)

const (
	Created Status = "created"
	Paid    Status = "paid"
)

const (
	PIX     = iota + 1 // 1
	DEBIT              // 2
	CREDIT             // 3
	INVOICE            // 4
)

var (
	ErrReachedLimit = errors.New("user reached the limit")
	ErrUnavaliable  = errors.New("item unavailable")
)
