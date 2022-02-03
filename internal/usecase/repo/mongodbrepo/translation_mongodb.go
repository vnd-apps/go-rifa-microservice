package mongodbrepo

import (
	"context"

	"github.com/evmartinelli/go-rifa-microservice/internal/entity"
	"github.com/evmartinelli/go-rifa-microservice/pkg/mongodb"
)

// TranslationRepo -.
type TranslationRepo struct {
	*mongodb.MongoDB
}

// New -.
func NewTranslation(mdb *mongodb.MongoDB) *TranslationRepo {
	return &TranslationRepo{mdb}
}

// GetHistory -.
func (r *TranslationRepo) GetHistory(ctx context.Context) ([]entity.Translation, error) {
	return nil, nil
}

// Store -.
func (r *TranslationRepo) Store(ctx context.Context, t entity.Translation) error {
	return nil
}
