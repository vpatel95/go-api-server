package main

import (
	"github.com/vpatel95/go-api-server/app"
)

func main() {

	app := &app.App{
		Host: "localhost",
		Port: 5000,
	}

	app.Run()
}
