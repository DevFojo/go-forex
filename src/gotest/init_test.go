package gotest

import (
	"github.com/MrFojo/go-forex/src/app"
	"github.com/MrFojo/go-forex/src/database"
	"testing"
)

func TestEnsureInitializeData(t *testing.T) { 
 
	app.EnsureInitializeData()

	a := database.Db.QueryRow("SELECT  COUNT(*) FROM rates")
	var count int
	if err := a.Scan(&count); err != nil { 
 
		t.Error(err.Error())
	}
	if count <= 0 { 
 
		t.Error("Date initialization saved no record")
	}
}
