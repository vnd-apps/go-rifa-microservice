package raffle_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	mock_raffle "github.com/evmartinelli/go-rifa-microservice/internal/core/mock/raffle"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	"github.com/evmartinelli/go-rifa-microservice/pkg/assert"
)

func ChangeStatusRaffleUseCase(t *testing.T) (*raffle.ChangeRaffleNumberStatusUseCase, *mock_raffle.MockRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := mock_raffle.NewMockRepo(mockCtl)

	changeRaffleNumberStatusUseCase := raffle.NewChangeRaffleNumberStatusUseCase(repo)

	return changeRaffleNumberStatusUseCase, repo
}

// generate a function to test change raffle number status use case.
func TestChangeRaffleNumberStatusUseCase(t *testing.T) {
	t.Parallel()

	changeRaffleNumberStatusUseCase, repo := ChangeStatusRaffleUseCase(t)

	// create a test struct
	tests := []test{
		{
			name: "Raffle Repo Error",
			mock: func() {
				repo.EXPECT().UpdateItem(context.Background(), "ak-valcan-test", 1).Return(errInternalServErr)
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
			err := changeRaffleNumberStatusUseCase.Run(context.Background(), "ak-valcan-test", 1)
			assert.Error(t, tc.err, err)
		})
	}
}
