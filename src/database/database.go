package database

import (
	"database/sql"

	"github.com/mrfojo/go-forex/src/config"
	"github.com/mrfojo/go-forex/src/utils"
)


var Db *sql.DB

func init() {
	db, err := sql.Open("sqlite3", config.DatabasePath)
	utils.ProcessError(err)
	Db = db
}
