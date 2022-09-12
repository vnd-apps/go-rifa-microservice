package restyhttp

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/skin"
)

const (
	_steamBaseURL = "https://steamcommunity.com/"
	_weaponType   = "Weapon"
	_steamCDNURL  = "http://cdn.steamcommunity.com/economy/image/"
)

// SteamWebAPI -.
type WebAPI struct {
	client *resty.Client
}

type Response struct {
	Descriptions []struct {
		MarketHash string `json:"market_hash_name"`
		Name       string `json:"name"`
		Marketable int    `json:"marketable"`
		Icon       string `json:"icon_url"`
		Tags       []struct {
			Category string `json:"category"`
		}
	}
}

// New -.
func NewSteamAPI() *WebAPI {
	client := resty.New()
	client.SetBaseURL(_steamBaseURL)

	return &WebAPI{
		client: client,
	}
}

// PlayerItens -.
func (s *WebAPI) PlayerItens(id string) (skin.Skin, error) {
	res := &Response{}
	skinItem := &skin.Skin{
		PlayerID: id,
	}

	err := getSteamInventory(s, res, id)
	if err != nil {
		return skin.Skin{}, err
	}

	createPlayerInventory(res, skinItem)

	if err != nil {
		return skin.Skin{}, err
	}

	return *skinItem, nil
}

func createPlayerInventory(res *Response, sk *skin.Skin) {
	for _, desc := range res.Descriptions {
		if desc.Marketable != 1 {
			continue
		}

		for _, tag := range desc.Tags {
			if tag.Category == _weaponType {
				sk.Items = append(sk.Items, skin.Item{
					Name:           desc.Name,
					MarketHashName: desc.MarketHash,
					Image:          _steamCDNURL + desc.Icon,
				})
			}
		}
	}
}

func getSteamInventory(s *WebAPI, res *Response, id string) error {
	resp, err := s.client.R().
		SetResult(&res).
		SetPathParams(map[string]string{"steam.id": id}).
		Get("/inventory/{steam.id}/730/2?l=en")
	if err != nil {
		return fmt.Errorf("TranslationWebAPI - Translate - trans.Translate: %w", err)
	}

	if resp.IsError() {
		return fmt.Errorf("TranslationWebAPI - Translate - trans.Translate: %w", err)
	}

	return nil
}
