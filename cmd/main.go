package main

import (
	"github.com/pashapdev/calc_go/internal/application"
)

// TODO: make logging
func main() {
	app := application.New()
	//app.Run()
	app.RunServer()
}
