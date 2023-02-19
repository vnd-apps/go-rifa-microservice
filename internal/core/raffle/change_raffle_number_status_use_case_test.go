package raffle_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	mock_raffle "github.com/evmartinelli/go-rifa-microservice/internal/core/mock/raffle"
	mock_shared "github.com/evmartinelli/go-rifa-microservice/internal/core/mock/shared"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/shared"
	"github.com/evmartinelli/go-rifa-microservice/pkg/assert"
)

func ChangeStatusRaffleUseCase(t *testing.T) (*raffle.ChangeRaffleNumberStatusUseCase, *mock_raffle.MockRepo, *mock_shared.MockAuth) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := mock_raffle.NewMockRepo(mockCtl)
	user := mock_shared.NewMockAuth(mockCtl)

	changeRaffleNumberStatusUseCase := raffle.NewChangeRaffleNumberStatusUseCase(repo, user)

	return changeRaffleNumberStatusUseCase, repo, user
}

// generate a function to test change raffle number status use case.
func TestChangeRaffleNumberStatusUseCase(t *testing.T) {
	t.Parallel()

	changeRaffleNumberStatusUseCase, repo, user := ChangeStatusRaffleUseCase(t)

	// create a test struct
	tests := []test{
		{
			name: "Raffle Repo Error",
			mock: func() {
				repo.EXPECT().UpdateItem(context.Background(), "ak-valcan-test", 1).Return(errInternalServErr)
				user.EXPECT().Claims(gomock.Any()).Return(&shared.User{}, nil)
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
