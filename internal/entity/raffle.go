// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Translation -.
type Raffle struct {
	Id           string
	Name         string `json:"name"  example:"en"`
	Status       string
	Value        int `json:"value"  example:"5"`
	TotalNumbers int `json:"totalNumbers"  example:"10"`
	TotalSold    string
}
