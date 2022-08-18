package raffle_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/evmartinelli/go-rifa-microservice/internal/entity"
	"github.com/evmartinelli/go-rifa-microservice/internal/usecase"
)

func raffle(t *testing.T) (*usecase.RaffleUseCase, *MockRaffleRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockRaffleRepo(mockCtl)

	raffle := usecase.NewRaffleUseCase(repo)

	return raffle, repo
}

func TestGetAvailable(t *testing.T) {
	t.Parallel()

	raffle, repo := raffle(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				repo.EXPECT().GetAvailableRaffle(context.Background()).Return(nil, nil)
			},
			res: []entity.Raffle(nil),
			err: nil,
		},
		{
			name: "result with error",
			mock: func() {
				repo.EXPECT().GetAvailableRaffle(context.Background()).Return(nil, errInternalServErr)
			},
			res: []entity.Raffle{},
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := raffle.GetAvailableRaffle(context.Background())

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestCreate(t *testing.T) {
	t.Parallel()

	raffle, repo := raffle(t)

	tests := []test{
		{
			name: "created",
			mock: func() {
				repo.EXPECT().Create(context.Background(), entity.Raffle{}).Return(nil)
			},
			err: nil,
		},
		{
			name: "repo error",
			mock: func() {
				repo.EXPECT().Create(context.Background(), entity.Raffle{}).Return(errInternalServErr)
			},
			res: nil,
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			err := raffle.Create(context.Background(), entity.Raffle{})

			require.ErrorIs(t, err, tc.err)
		})
	}
}
