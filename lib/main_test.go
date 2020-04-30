package lib_test

import (
	"cardcaptor/lib"
	"database/sql"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

const (
	dbPath = "../res/hearthstone.db"
)

func TestMain(m *testing.M) {
	println("+=+=+ Start Test +=+=+")
	if !lib.IsFileExist(dbPath) {
		lib.CreateNewDB(dbPath)
	}
	db, _ = sql.Open("sqlite3", dbPath)
	defer db.Close()
	os.Exit(m.Run())
}
