package skin_test

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	"github.com/golang/mock/gomock"

// 	mock_skin "github.com/evmartinelli/go-rifa-microservice/internal/core/mock/mock_skin"
// 	"github.com/evmartinelli/go-rifa-microservice/internal/core/skin"
// 	"github.com/evmartinelli/go-rifa-microservice/pkg/assert"
// )

// type test struct {
// 	name string
// 	mock func()
// 	res  interface{}
// 	err  error
// }

// var errInternalServErr = errors.New("internal server error")

// func generatePlayerInventory(t *testing.T) (*skin.PlayerInventoryUseCase, *mock_skin.MockPlayerSkinRepo, *mock_skin.MockSteamWebAPI) {
// 	t.Helper()

// 	mockCtl := gomock.NewController(t)
// 	defer mockCtl.Finish()

// 	repo := mock_skin.NewMockPlayerSkinRepo(mockCtl)
// 	webapi := mock_skin.NewMockSteamWebAPI(mockCtl)

// 	playerInventoryUseCase := skin.NewPlayerInventoryUseCase(repo, webapi)

// 	return playerInventoryUseCase, repo, webapi
// }

// func TestHistory(t *testing.T) {
// 	t.Parallel()

// 	generatePlayerInventory, repo, webAPI := generatePlayerInventory(t)

// 	tests := []test{
// 		{
// 			name: "empty result",
// 			mock: func() {
// 				webAPI.EXPECT().PlayerItens(gomock.Any()).Return(skin.Skin{}, nil)
// 				repo.EXPECT().Create(context.Background(), skin.Skin{}).Return(nil)
// 			},
// 			res: skin.Skin{},
// 			err: nil,
// 		},
// 		{
// 			name: "web API error",
// 			mock: func() {
// 				webAPI.EXPECT().PlayerItens(gomock.Any()).Return(skin.Skin{}, errInternalServErr)
// 			},
// 			res: skin.Skin{},
// 			err: errInternalServErr,
// 		},
// 		{
// 			name: "repo error",
// 			mock: func() {
// 				webAPI.EXPECT().PlayerItens(gomock.Any()).Return(skin.Skin{}, nil)
// 				repo.EXPECT().Create(context.Background(), skin.Skin{}).Return(errInternalServErr)
// 			},
// 			res: skin.Skin{},
// 			err: errInternalServErr,
// 		},
// 	}

// 	for _, tc := range tests {
// 		tc := tc

// 		t.Run(tc.name, func(t *testing.T) {
// 			t.Parallel()

// 			tc.mock()

// 			res, err := generatePlayerInventory.Run(context.Background(), "ID")
// 			assert.DeepEqual(t, tc.res, res)
// 			assert.Error(t, tc.err, err)
// 		})
// 	}
// }
