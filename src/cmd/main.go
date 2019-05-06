package main

import (
	"github.com/MrFojo/go-forex/src/app"
)

func main() {
	app.EnsureInitializeData()
	app.Run()
}
