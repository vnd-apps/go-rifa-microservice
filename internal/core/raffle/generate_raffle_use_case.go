package raffle

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/shared"
)

type GenerateRaffleUseCase struct {
	repo Repo
	uuid shared.UUIDGenerator
	slug shared.SlugGenerator
}

func NewGenerateRaffleUseCase(r Repo, u shared.UUIDGenerator, s shared.SlugGenerator) *GenerateRaffleUseCase {
	return &GenerateRaffleUseCase{
		repo: r,
		uuid: u,
		slug: s,
	}
}

func (uc *GenerateRaffleUseCase) Run(ctx context.Context, model *Raffle) error {
	model.ID = uc.uuid.Generate()
	model.Slug = uc.slug.Generate(model.Name)
	model.Status = Open

	for i := 1; i <= model.Quantity; i++ {
		a := Variation{
			ID:     model.ID,
			Number: i,
			Name:   "Number",
			Status: Available,
		}
		model.Variation = append(model.Variation, a)
	}

	return uc.repo.Create(ctx, model)
}
