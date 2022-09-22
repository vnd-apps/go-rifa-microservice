package order

type Request struct {
	ID            string
	ProductID     string
	Total         int
	PaymentMethod PaymentMethod
	Items         []int
}

type Order struct {
	ID            string
	ProductID     string
	UserID        string
	Total         int
	PaymentMethod PaymentMethod
	Items         []int
	Pix           Pix
}

type Pix struct {
	ID           int
	Status       string
	QRCode       string
	QRCodeBase64 string
}

type PaymentMethod string

const (
	PIX = "pix"
)
