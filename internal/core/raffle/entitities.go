// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package raffle

type Raffle struct {
	ID           string    `json:"id"  example:"61f0c143ad06223fa03910b0"`
	Name         string `json:"name"  example:"Rifa"`
	Status       string `json:"status"  example:"Available"`
	Value        int    `json:"value"  example:"5"`
	TotalNumbers int    `json:"totalNumbers"  example:"10"`
	TotalSold    int    `json:"totalSold"  example:"1"`
}
