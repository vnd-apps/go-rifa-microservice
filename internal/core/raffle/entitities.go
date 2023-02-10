// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package raffle

type Request struct {
	Name        string  `json:"name" binding:"required" example:"Rifa"`
	Description string  `json:"description" binding:"required" example:"Rifa"`
	ImageURL    string  `json:"imageURL" binding:"required" example:"1"`
	UnitPrice   float32 `json:"unitPrice" binding:"required" example:"5"`
	Quantity    int     `json:"quantity" binding:"required" example:"10"`
}

type Raffle struct {
	ID              string    `json:"id"  example:"61f0c143ad06223fa03910b0"`
	Name            string    `json:"name"  example:"Rifa"`
	Description     string    `json:"description"  example:"Rifa description"`
	Slug            string    `json:"slug"  example:"butterfly-32"`
	Status          Status    `json:"status"  example:"open"`
	ImageURL        string    `json:"imageURL"  example:"1"`
	UnitPrice       float32   `json:"unitPrice"  example:"5"`
	UserLimit       int       `json:"userLimit,omitempty"  example:"10"`
	Quantity        int       `json:"quantity"  example:"10"`
	PrizeDrawNumber int       `json:"PrizeDrawNumber,omitempty"  example:"10"`
	Numbers         []Numbers `json:"numbers,omitempty"`
}

type Numbers struct {
	ID     string     `json:"id"  example:"61f0c143ad06223fa03910b0"`
	Number int        `json:"number"  example:"5"`
	Slug   string     `json:"name"  example:"Number"`
	Status ItemStatus `json:"status"  example:"paid"`
}

type (
	ItemStatus string
	Status     string
)

const (
	Paid      ItemStatus = "paid"
	Pending   ItemStatus = "pending"
	Available ItemStatus = "available"
	Open      Status     = "open"
	Cloed     Status     = "closed"
)
