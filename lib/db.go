package lib

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"log"
	"strings"

	// _
	_ "github.com/mattn/go-sqlite3"
)

// newConnection is
func newConnection(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// CreateNewDB is
func CreateNewDB(path string) {
	if IsFileExist(path) {
		log.Fatal(errors.New("file is already exists"))
	}
	conn := newConnection(path)
	defer conn.Close()

	log.Println("Initialize database...")
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
