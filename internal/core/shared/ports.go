package shared

type UUIDGenerator interface {
	Generate() string
}
type SlugGenerator interface {
	Generate(text string) string
}
