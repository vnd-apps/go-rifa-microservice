package slug

import "github.com/gosimple/slug"

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Generate(text string) string {
	return slug.Make(text)
}
