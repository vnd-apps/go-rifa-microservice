package raffle_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	mock_raffle "github.com/evmartinelli/go-rifa-microservice/internal/core/mock/mock_raffle"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	"github.com/evmartinelli/go-rifa-microservice/pkg/assert"
)

func listRaffleUseCase(t *testing.T) (*raffle.ListRaffleUseCase, *mock_raffle.MockRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := mock_raffle.NewMockRepo(mockCtl)

	listRaffleUseCase := raffle.NewListRaffleUseCase(repo)

	return listRaffleUseCase, repo
}

func TestList(t *testing.T) {
	t.Parallel()

	listRaffleUseCase, repo := listRaffleUseCase(t)

	tests := []test{
		{
			name: "Empty List Raffle",
			mock: func() {
				repo.EXPECT().GetAll(context.Background()).Return([]raffle.Raffle{}, nil)
			},
			err: nil,
			res: []raffle.Raffle{},
		},
		{
			name: "Raffle Repo Error",
			mock: func() {
				repo.EXPECT().GetAll(context.Background()).Return(nil, errInternalServErr)
			},
			err: errInternalServErr,
			res: []raffle.Raffle{},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()
			res, err := listRaffleUseCase.Run(context.Background())
			assert.DeepEqual(t, tc.res, res)
			assert.Error(t, tc.err, err)
		})
	}
}
