package main

import (
	"github.com/devFojo/go-forex/app"
)

func main() {
	app.EnsureInitializeData()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
