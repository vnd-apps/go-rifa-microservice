package pix

import (
	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/pix/fake"
)

type Payment struct{}

func NewPixPayment() *fake.Pix {
	return fake.NewFakePixPayment()
}
