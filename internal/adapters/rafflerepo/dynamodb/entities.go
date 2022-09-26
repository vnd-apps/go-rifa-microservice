package dynamodb

type DynamoRaffle struct {
	PK           string
	SK           string
	GSI1PK       string
	ID           string
	Name         string
	Description  string
	Slug         string
	Status       string
	ImageURL     string
	ItemType     string
	UnitPrice    int
	Quantity     int
	UserLimit    int
	SortedNumber int
}

type DynamoRaffleItem struct {
	PK       string
	SK       string
	ID       string
	Number   int
	Name     string
	Status   string
	ItemType string
}

type DynamoRecRaffle struct {
	PK           string
	SK           string
	GSI1PK       string
	ID           string
	Name         string
	Description  string
	Slug         string
	Status       string
	ImageURL     string
	ItemType     string
	UnitPrice    int
	Quantity     int
	UserLimit    int
	SortedNumber int
	Number       int
}
