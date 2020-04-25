package lib

import (
	"database/sql"
	"io/ioutil"
	"log"
	"strings"

	// _
	_ "github.com/mattn/go-sqlite3"
)

// NewConnection is
func NewConnection(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// CreateNewDB is
func CreateNewDB(path string) {
	conn := NewConnection(path)
	createScheme(conn)
}

func createScheme(db *sql.DB) {
	contents, err := ioutil.ReadFile(schemeSQL)
	if err != nil {
		log.Fatal(err)
	}

	sqls := strings.Split(string(contents), ";\n")

	for _, sql := range sqls {
		db.Exec(sql)
	}

}
