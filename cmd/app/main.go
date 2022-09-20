package main

import (
	"github.com/evmartinelli/go-rifa-microservice/internal/app"
)

func main() {
	c := app.NewContext()
	c.Config()
	c.HTTPServer()
}
