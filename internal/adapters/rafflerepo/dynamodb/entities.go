package dynamodb

type DynamoRaffle struct {
	PK          string
	SK          string
	GSI1PK      string
	ID          string
	Name        string
	Description string
	Slug        string
	Status      string
	ImageURL    string
	UnitPrice   int
	Quantity    int
}

type DynamoRaffleItem struct {
	PK     string
	SK     string
	ID     string
	Number int
	Name   string
	Status string
}
