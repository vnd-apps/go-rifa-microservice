package pix

import "github.com/evmartinelli/go-rifa-microservice/internal/adapters/pix/fake"

func NewPixPayment() *fake.Pix {
	return fake.NewFakePixPayment()
}
