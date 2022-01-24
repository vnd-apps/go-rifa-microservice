// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Translation -.
type Raffle struct {
	Id           string
	Name         string   `json:"name"  example:"en"`
	Images       []string `json:"images"     example:"текст для перевода"`
	Status       string
	Value        string `json:"value"  example:"text for translation"`
	TotalNumbers string `json:"totalNumbers"  example:"text for translation"`
	TotalSold    string
}
