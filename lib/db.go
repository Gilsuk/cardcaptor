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

// NewDB is
func NewDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// CreateNewDB is
func CreateNewDB(path string) error {
	if IsFileExist(path) {
		log.Fatal(errors.New("file is already exists"))
	}
	db := NewDB(path)
	defer db.Close()

	log.Println("Initialize database...")
	err := createScheme(db)

	if err != nil {
		return err
	}
	return nil
}

func createScheme(db *sql.DB) error {
	contents, err := ioutil.ReadFile(schemeSQL)
	if err != nil {
		return err
	}

	sqls := strings.Split(string(contents), ";\n")

	for _, sql := range sqls {
		if _, err = db.Exec(sql); err != nil {
			return err
		}
	}

	return nil
}

// CardIDByOffset is
func CardIDByOffset(db *sql.DB, offset int) (id int, err error) {
	query := `
		SELECT card FROM card
		ORDER BY card
		LIMIT 1 OFFSET ?
	`
	row := db.QueryRow(query, offset)
	err = row.Scan(&id)
	return
}

// CardsCount is
func CardsCount(db *sql.DB) (count int, err error) {
	query := `
		SELECT count(*) FROM card
	`
	row := db.QueryRow(query)
	err = row.Scan(&count)
	return
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
		INSERT INTO class (class, slug, name, card)
		VALUES (?, ?, ?, ?)
		`
	stmt, err := db.Prepare(query)
	defer stmt.Close()
	_, err = stmt.Exec(c.ID, c.Slug, c.Name, c.CardID)
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

// Insert must be called after sets are completly inserted
func (s *SetGroup) Insert(db *sql.DB) error {

	query := `
		INSERT INTO setgroup (slug, year, name, standard)
		VALUES (?, ?, ?, ?)
		`
	mappingQuery := `
		INSERT INTO groups (setGroup, cardSet)
		VALUES (?, (
			SELECT cardset FROM cardset
			WHERE slug = ?
		))
		`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(s.Slug, s.Year, s.Name, s.Standard)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()

	mappingStmt, err := db.Prepare(mappingQuery)
	if err != nil {
		return err
	}
	defer mappingStmt.Close()

	if s.Sets == nil {
		return nil
	}

	for _, set := range s.Sets {
		mappingStmt.Exec(id, set)
	}

	return nil
}

// Delete is
func (s *SetGroup) Delete(db *sql.DB) error {
	return errors.New("SetGroup.Delete()is currently not implemented")
}

// Insert is
func (s *Set) Insert(db *sql.DB) error {
	query := `
		INSERT INTO cardset (cardset, name, slug, releasedate, type)
		VALUES (?, ?, ?, ?, ?)
		`
	stmt, _ := db.Prepare(query)
	defer stmt.Close()
	_, err := stmt.Exec(s.ID, s.Name, s.Slug, s.ReleaseDate, s.Type)
	return err
}

// Delete is
func (s *Set) Delete(db *sql.DB) error {
	return errors.New("Set.Delete()is currently not implemented")
}

// Insert must be called after all cards was inserted
func (s *Arena) Insert(db *sql.DB) error {
	query := `
		UPDATE card SET arena = 1
		WHERE card = ?
		`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(*s))
	if err != nil {
		return err
	}
	return nil
}

// Delete is
func (s *Arena) Delete(db *sql.DB) error {
	query := `
		UPDATE card SET arena = 0
		WHERE card = ?
		`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(*s))
	if err != nil {
		return err
	}
	return nil
}

// Insert is
func (c *Card) Insert(db *sql.DB) (err error) {
	if err = c.insertBasicInfo(db); err != nil {
		return
	}
	if err = c.insertClasses(db); err != nil {
		return
	}
	if err = c.insertMechanism(db); err != nil {
		return
	}
	if err = c.insertFamily(db); err != nil {
		return
	}
	return
}

// InsertMissingData is
func InsertMissingData(db *sql.DB) error {
	keyword := Keyword{
		ID:   7,
		Slug: "enraged",
		Name: "분노",
		Ref:  "피해를 받은 상태면 공격력을 얻습니다.",
		Text: "피해를 받은 상태면 공격력을 {0}얻습니다.",
	}

	if err := keyword.Insert(db); err != nil {
		return err
	}

	etc := Type{
		ID:   6,
		Name: "기타",
		Slug: "etc",
	}

	hp := Type{
		ID:   6,
		Name: "영웅 능력",
		Slug: "hero-power",
	}

	if err := etc.Insert(db); err != nil {
		return err
	}
	if err := hp.Insert(db); err != nil {
		return err
	}

	return nil
}

// VacuumDB delete unused rows
func VacuumDB(db *sql.DB) (err error) {
	query := `
		DELETE FROM family
		WHERE NOT EXISTS (
			SELECT card FROM card
			WHERE family.child = card.card
		)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return
	}

	db.Exec("vacuum")

	return
}

func (c *Card) insertFamily(db *sql.DB) (err error) {

	if c.Child == nil || len(c.Child) == 0 {
		return
	}

	query := `
		INSERT INTO family (
			parent, child
		) VALUES (
			?, ?
		)
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	for _, cid := range c.Child {
		_, err = stmt.Exec(c.Card, cid)
		if err != nil {
			return
		}
	}

	return
}

func (c *Card) insertClasses(db *sql.DB) error {
	query := `
		INSERT INTO classes (
			card, class
		) VALUES (
			?, ?
		)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, cid := range c.MultiClassIDs {
		_, err = stmt.Exec(c.Card, cid)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Card) insertMechanism(db *sql.DB) error {
	query := `
		INSERT INTO mechanism (
			card, keyword
		) VALUES (
			?, ?
		)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, kid := range c.Keyword {
		_, err = stmt.Exec(c.Card, kid)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Card) insertBasicInfo(db *sql.DB) error {
	query := `
		INSERT INTO card (
			card, slug, class, type, cardset, rarity, race, artist,
			name, text, flavor, img, cropimg, cost, health, attack,
			armor, collectable, durability
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?
		)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		c.Card, c.Slug, c.Class, c.Type, c.Set, c.Rarity, c.Race, c.Artist,
		c.Name, c.Text, c.Flavor, c.Img, c.CropImg, c.Cost, c.Health, c.Attack,
		c.Armor, c.Collectible, c.Durability,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete is
func (c *Card) Delete(db *sql.DB) error {
	return errors.New("Card.Delete()is currently not implemented")
}

// Insert is insert metadata without arena info
// arenaids must be inserted after all cards data were inserted
func (m *Meta) Insert(db *sql.DB) error {
	for _, v := range m.Classes {
		if err := v.Insert(db); err != nil {
			return err
		}
	}
	for _, v := range m.Keywords {
		if err := v.Insert(db); err != nil {
			return err
		}
	}
	for _, v := range m.Races {
		if err := v.Insert(db); err != nil {
			return err
		}
	}
	for _, v := range m.Rarities {
		if err := v.Insert(db); err != nil {
			return err
		}
	}
	for _, v := range m.Types {
		if err := v.Insert(db); err != nil {
			return err
		}
	}
	for _, v := range m.Sets {
		if err := v.Insert(db); err != nil {
			return err
		}
	}
	for _, v := range m.SetGroups {
		if err := v.Insert(db); err != nil {
			return err
		}
	}
	return nil
}

// InsertArena is
func (m *Meta) InsertArena(db *sql.DB) error {
	for _, v := range m.Arenas {
		if err := v.Insert(db); err != nil {
			return err
		}
	}

	return nil
}

// ClassCards is
func (m *Meta) ClassCards() (ids []int) {
	for _, v := range m.Classes {
		if v.CardID != 0 {
			ids = append(ids, v.CardID)
		}
	}
	return
}
