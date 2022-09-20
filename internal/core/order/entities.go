package order

type PlaceOrderRequest struct {
	ID            string
	ProductID     string
	Total         int
	PaymentMethod PaymentMethod
	Items         []Item
}

type PlaceOrderResponse struct {
	ID            string
	ProductID     string
	Total         int
	PaymentMethod PaymentMethod
	Items         []Item
	Pix           Pix
}

type Pix struct {
	ID           int
	Status       string
	QRCode       string
	QRCodeBase64 string
}

type Item struct {
	ID     string
	Number int
	Price  int
}

type PaymentMethod string

const (
	PIX = "pix"
)
