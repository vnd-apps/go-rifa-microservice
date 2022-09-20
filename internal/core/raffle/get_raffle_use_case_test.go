package raffle_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	"github.com/evmartinelli/go-rifa-microservice/pkg/assert"
)

func getRaffleUseCase(t *testing.T) (*raffle.GetRaffleUseCase, *MockRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockRepo(mockCtl)

	getRaffleUseCase := raffle.NewGetRaffleUseCase(repo)

	return getRaffleUseCase, repo
}

func TestGet(t *testing.T) {
	t.Parallel()

	getRaffleUseCase, repo := getRaffleUseCase(t)

	tests := []test{
		{
			name: "Empty List Raffle",
			mock: func() {
				repo.EXPECT().GetProduct(context.Background(), "b615df9c-dfd6-4709-9d30-10552057efde").Return(raffle.Raffle{}, nil)
			},
			err: nil,
			res: raffle.Raffle{},
		},
		{
			name: "Raffle Repo Error",
			mock: func() {
				repo.EXPECT().GetProduct(context.Background(), "b615df9c-dfd6-4709-9d30-10552057efde").Return(raffle.Raffle{}, errInternalServErr)
			},
			err: errInternalServErr,
			res: raffle.Raffle{},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()
			res, err := getRaffleUseCase.Run(context.Background(), "b615df9c-dfd6-4709-9d30-10552057efde")
			assert.DeepEqual(t, tc.res, res)
			assert.Error(t, tc.err, err)
		})
	}
}
