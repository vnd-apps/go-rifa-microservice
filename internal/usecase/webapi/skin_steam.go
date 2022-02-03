package webapi

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/evmartinelli/go-rifa-microservice/internal/entity"
)

const (
	_steamBaseURL = "https://steamcommunity.com/"
	_weaponType   = "Weapon"
)

// SteamWebAPI -.
type SteamWebAPI struct {
	client *resty.Client
}

type Response struct {
	Descriptions []struct {
		MarketHash string `json:"market_hash_name"`
		Name       string `json:"name"`
		Tags       []struct {
			Category string `json:"category"`
		}
	}
}

// New -.
func NewSteamAPI() *SteamWebAPI {
	client := resty.New()
	client.SetBaseURL(_steamBaseURL)

	return &SteamWebAPI{
		client: client,
	}
}

// Translate -.
func (s *SteamWebAPI) PlayerItens(translation entity.Translation, id string) (entity.Skin, error) {
	res := &Response{}
	skin := &entity.Skin{
		PlayerID: id,
	}

	_, err := getSteamInventory(s, res, id)
	if err != nil {
		return entity.Skin{}, err
	}

	createPlayerInventory(res, skin)

	return *skin, nil
}

func createPlayerInventory(res *Response, skin *entity.Skin) {
	for _, desc := range res.Descriptions {
		for _, tag := range desc.Tags {
			if tag.Category == _weaponType {
				skin.Items = append(skin.Items, entity.Item{
					Name:           desc.Name,
					MarketHashName: desc.MarketHash,
				})
			}
		}
	}
}

func getSteamInventory(s *SteamWebAPI, res *Response, id string) (entity.Skin, error) {
	resp, err := s.client.R().
		SetResult(&res).
		SetPathParams(map[string]string{"steam.id": id}).
		Get("/inventory/{steam.id}/730/2?l=en")
	if err != nil {
		return entity.Skin{}, fmt.Errorf("TranslationWebAPI - Translate - trans.Translate: %w", err)
	}

	if resp.IsError() {
		return entity.Skin{}, fmt.Errorf("TranslationWebAPI - Translate - trans.Translate: %w", err)
	}

	return entity.Skin{}, nil
}
