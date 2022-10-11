package raffle_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/evmartinelli/go-rifa-microservice/internal/adapters/idgenerator"
	mock_raffle "github.com/evmartinelli/go-rifa-microservice/internal/core/mock/mock_raffle"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	"github.com/evmartinelli/go-rifa-microservice/pkg/assert"
)

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

var errInternalServErr = errors.New("internal server error")

func generateRaffleUseCase(t *testing.T) (*raffle.GenerateRaffleUseCase, *mock_raffle.MockRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := mock_raffle.NewMockRepo(mockCtl)
	uuid := idgenerator.NewUUIDGenerator()
	slug := idgenerator.NewSlugGenerator()

	generateRaffleUseCase := raffle.NewGenerateRaffleUseCase(repo, uuid, slug)

	return generateRaffleUseCase, repo
}

func TestGenerate(t *testing.T) {
	t.Parallel()

	raffleUseCase, repo := generateRaffleUseCase(t)

	tests := []test{
		{
			name: "Create Raffle",
			mock: func() {
				repo.EXPECT().Create(context.Background(), gomock.Any()).Return(nil)
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()
			err := raffleUseCase.Run(context.Background(), &raffle.Raffle{Quantity: 5})
			assert.Error(t, tc.err, err)
		})
	}
}
