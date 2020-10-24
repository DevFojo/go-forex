package database

import (
	"database/sql"

	"github.com/devFojo/go-forex/config"
	"github.com/devFojo/go-forex/utils"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func init() {
	db, err := sql.Open("sqlite3", config.DatabasePath)
	utils.ProcessError(err)
	Db = db
}
