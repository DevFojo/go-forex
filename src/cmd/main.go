package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/mrfojo/go-forex/src/app"
)

func main() {
	app.InitializeData()
}
