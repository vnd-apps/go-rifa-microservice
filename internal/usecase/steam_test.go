package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/evmartinelli/go-rifa-microservice/internal/entity"
	"github.com/evmartinelli/go-rifa-microservice/internal/usecase"
)

func steam(t *testing.T) (*usecase.SteamUseCase, *MockPlayerSkinRepo, *MockSteamWebAPI) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockPlayerSkinRepo(mockCtl)
	webAPI := NewMockSteamWebAPI(mockCtl)

	steam := usecase.NewSteam(repo, webAPI)

	return steam, repo, webAPI
}

func TestCreateSkin(t *testing.T) {
	t.Parallel()

	steam, repo, webAPI := steam(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				webAPI.EXPECT().PlayerItens("123").Return(entity.Skin{}, nil)
				repo.EXPECT().Create(context.Background(), entity.Skin{}).Return(nil)
			},
			res: entity.Skin{},
			err: nil,
		},
		{
			name: "web API error",
			mock: func() {
				webAPI.EXPECT().PlayerItens("123").Return(entity.Skin{}, errInternalServErr)
			},
			res: entity.Skin{},
			err: errInternalServErr,
		},
		{
			name: "repo error",
			mock: func() {
				webAPI.EXPECT().PlayerItens("123").Return(entity.Skin{}, nil)
				repo.EXPECT().Create(context.Background(), entity.Skin{}).Return(errInternalServErr)
			},
			res: entity.Skin{},
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := steam.GetPlayerInventory(context.Background(), "123")

			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
