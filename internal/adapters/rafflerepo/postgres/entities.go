package postgres

type Raffle struct {
	Name            string
	Description     string
	Slug            string
	Status          string
	ImageURL        string
	UnitPrice       float32 `sql:"type:decimal(10,2);"`
	UserLimit       int
	Quantity        int
	PrizeDrawNumber int
}

type RaffleNumbers struct {
	Slug   string
	Number int
	Status string
}
