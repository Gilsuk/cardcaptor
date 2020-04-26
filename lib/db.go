package lib

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"log"
	"strings"
)

// DBItem is
type DBItem interface {
	Insert(db *sql.DB) error
	Delete(db *sql.DB) error
}

// newDB is
func newDB(path string) *sql.DB {
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
	db := newDB(path)
	defer db.Close()

	log.Println("Initialize database...")
	createScheme(db)
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

// Insert is
func (k *Keyword) Insert(db *sql.DB) error {
	query := `
		INSERT INTO keyword (keyword, slug, name, ref, text)
		VALUES (?, ?, ?, ?, ?)
		`
	stmt, _ := db.Prepare(query)
	defer stmt.Close()
	_, err := stmt.Exec(k.ID, k.Slug, k.Name, k.Ref, k.Text)
	return err
}

// Delete is
func (k *Keyword) Delete(db *sql.DB) error {
	query := `
		DELETE FROM keyword
		WHERE keyword = ?
		`
	stmt, _ := db.Prepare(query)
	defer stmt.Close()
	_, err := stmt.Exec(k.ID)
	return err
}
