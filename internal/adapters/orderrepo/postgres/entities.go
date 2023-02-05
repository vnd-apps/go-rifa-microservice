package postgres

type Order struct {
	ID            int64
	ProductID     string
	UserID        string
	Total         float32
	PaymentMethod string
}

type OrderItems struct {
	OrderID int64
	Numbers int
}

type PixPayment struct {
	OrderID      int64
	Status       string
	QRCode       string
	QRCodeBase64 string
}
