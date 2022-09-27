package dynamodb

type DynamoOrder struct {
	PK            string
	SK            string
	GSI1PK        string
	ID            string
	ProductID     string
	UserID        string
	Total         int
	PaymentMethod string
	Items         []int
	Pix           struct {
		ID           int
		Status       string
		QRCode       string
		QRCodeBase64 string
	}
	ItemType string
}
