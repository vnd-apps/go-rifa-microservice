package steam

import "github.com/evmartinelli/go-rifa-microservice/internal/adapters/steam/restyhttp"

func NewSteamAPI() *restyhttp.WebAPI {
	return restyhttp.NewSteamAPI()
}
