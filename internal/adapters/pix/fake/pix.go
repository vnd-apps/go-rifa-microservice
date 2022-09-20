package fake

import "github.com/evmartinelli/go-rifa-microservice/internal/core/order"

type Pix struct {
	BaseURL string
	State   string
	Code    string
}

func NewFakePixPayment() *Pix {
	return &Pix{}
}

func (p *Pix) GeneratePix() (order.Pix, error) {
	return order.Pix{
		ID:           54323453,
		Status:       "pending",
		QRCodeBase64: "iVBORw0KGgoAAAANSUhEUgAABRQAAAUUCAYAAACu5p7oAAAABGdBTUEAALGPC",
		QRCode:       "00020126600014br.gov.bcb.pix0117john@yourdomain.com0217additional data520400005303986540510.005802BR5913Maria Silva6008Brasilia62070503***6304E2CA",
	}, nil
}
