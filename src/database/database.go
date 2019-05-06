package database

import (
	"database/sql"

	"github.com/MrFojo/go-forex/src/config"
	"github.com/MrFojo/go-forex/src/utils"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func init() {
	db, err := sql.Open("sqlite3", config.DatabasePath)
	utils.ProcessError(err)
	Db = db
}
