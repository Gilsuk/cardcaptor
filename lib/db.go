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

// Insert is
func (t *Type) Insert(db *sql.DB) error {
	query := `
		INSERT INTO type (type, slug, name)
		VALUES (?, ?, ?)
		`
	stmt, _ := db.Prepare(query)
	defer stmt.Close()
	_, err := stmt.Exec(t.ID, t.Slug, t.Name)
	return err
}

// Delete is
func (t *Type) Delete(db *sql.DB) error {
	query := `
		DELETE FROM type
		WHERE type = ?
		`
	stmt, _ := db.Prepare(query)
	defer stmt.Close()
	_, err := stmt.Exec(t.ID)
	return err
}

// Insert is
func (r *Race) Insert(db *sql.DB) error {
	query := `
		INSERT INTO race (race, slug, name)
		VALUES (?, ?, ?)
		`
	stmt, _ := db.Prepare(query)
	defer stmt.Close()
	_, err := stmt.Exec(r.ID, r.Slug, r.Name)
	return err
}

// Delete is
func (r *Race) Delete(db *sql.DB) error {
	query := `
		DELETE FROM race
		WHERE race = ?
		`
	stmt, _ := db.Prepare(query)
	defer stmt.Close()
	_, err := stmt.Exec(r.ID)
	return err
}

// Insert is
func (c *Class) Insert(db *sql.DB) error {
	query := `
		INSERT INTO class (class, slug, name)
		VALUES (?, ?, ?)
		`
	stmt, _ := db.Prepare(query)
	defer stmt.Close()
	_, err := stmt.Exec(c.ID, c.Slug, c.Name)
	return err
}

// Delete is
func (c *Class) Delete(db *sql.DB) error {
	query := `
		DELETE FROM class
		WHERE class = ?
		`
	stmt, _ := db.Prepare(query)
	defer stmt.Close()
	_, err := stmt.Exec(c.ID)
	return err
}

// Insert is
func (r *Rarity) Insert(db *sql.DB) error {
	query := `
		INSERT INTO rarity (rarity, slug, name)
		VALUES (?, ?, ?)
		`
	stmt, _ := db.Prepare(query)
	defer stmt.Close()
	_, err := stmt.Exec(r.ID, r.Slug, r.Name)
	return err
}

// Delete is
func (r *Rarity) Delete(db *sql.DB) error {
	query := `
		DELETE FROM rarity
		WHERE rarity = ?
		`
	stmt, _ := db.Prepare(query)
	defer stmt.Close()
	_, err := stmt.Exec(r.ID)
	return err
}
