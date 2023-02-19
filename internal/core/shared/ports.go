package shared

//go:generate mockgen -source=ports.go -destination=../mock/shared/mock_shared.go
type UUIDGenerator interface {
	Generate() string
}
type SlugGenerator interface {
	Generate(text string) string
}

type Auth interface {
	ExtractToken(bearerToken string) string
	CheckIsValid(bearerToken string) error
	Claims(bearerToken string) (*User, error)
}
