package postgres

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	RaffleSlug    string
	UserID        string
	Total         float32
	PaymentMethod int
	Status        string
}

type OrderItems struct {
	gorm.Model
	OrderID uint
	Number  int
}

type PixPayment struct {
	gorm.Model
	OrderID      uint
	Status       string
	QRCode       string
	QRCodeBase64 string
}
