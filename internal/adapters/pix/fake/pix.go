package fake

import (
	"fmt"

	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/idgenerator/uuid"
)

type Pix struct {
	BaseURL string
	State   string
	Code    string
}

func NewFakePixPayment() *Pix {
	return &Pix{}
}

func (p *Pix) GeneratePix(state string) string {
	p.State = state
	p.Code = uuid.NewGenerator().Generate()

	return fmt.Sprintf("%s/login/github/confirm?state=%s&code=%s", p.BaseURL, p.State, p.Code)
}
