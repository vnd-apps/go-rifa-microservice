// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package skin

// Translation -.
type Skin struct {
	ID       string `json:"id,omitempty"`
	PlayerID string `json:"steam_id"`
	Items    []Item `json:"items"`
}

type Item struct {
	Name           string `json:"name"`
	Image          string `json:"icon_url"`
	MarketHashName string `json:"market_hash_name"`
}
