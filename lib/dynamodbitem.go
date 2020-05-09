package lib

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"
)

// CardItem is
type CardItem struct {
	ID          int
	Slug        string
	Class       string
	Classes     []string
	CardType    string
	CardSet     string
	Rarity      string
	Race        string
	Keywords    []KeywordItem
	Artist      string
	Name        string
	Text        string
	Flavor      string
	Img         string
	CropImg     string
	Cost        int
	Health      int
	Attack      int
	Armor       int
	Durability  int
	Arena       bool
	Collectible bool
	Standard    bool
	Parents     []CardItem
	Children    []CardItem
}

// KeywordItem is
type KeywordItem struct {
	Name string
	Text string
}

// Export is
func (c *CardItem) Export(dir string) {
	os.MkdirAll(dir, os.ModePerm)

	bytes, _ := json.Marshal(*c)
	ioutil.WriteFile(dir+"/"+c.Slug+".json", bytes, os.ModePerm)
}

// NewCardItem is
func NewCardItem(db *sql.DB, id int) (c CardItem, err error) {
	if err = c.fillBasics(db, id); err != nil {
		return
	}
	if err = c.fillClasses(db); err != nil {
		return
	}
	if err = c.fillKeywords(db); err != nil {
		return
	}
	if err = c.fillParents(db); err != nil {
		return
	}
	if err = c.fillChildren(db); err != nil {
		return
	}
	return
}

func newInnerCardItem(db *sql.DB, id int) (c CardItem, err error) {
	if err = c.fillBasics(db, id); err != nil {
		return
	}
	if err = c.fillClasses(db); err != nil {
		return
	}
	if err = c.fillKeywords(db); err != nil {
		return
	}
	return
}

func (c *CardItem) fillBasics(db *sql.DB, id int) (err error) {
	query := `
		SELECT 
			id,slug,name,class,type,
			rarity,race,cost,health,attack,armor,
			durability,arena,collectable,text,flavor,
			artist,cardSet,img,cropImg,standard
		FROM vCard
		WHERE id = ?
	`
	row := db.QueryRow(query, id)

	err = row.Scan(&c.ID, &c.Slug, &c.Name, &c.Class, &c.CardType,
		&c.Rarity, &c.Race, &c.Cost, &c.Health, &c.Attack, &c.Armor,
		&c.Durability, &c.Arena, &c.Collectible, &c.Text, &c.Flavor,
		&c.Artist, &c.CardSet, &c.Img, &c.CropImg, &c.Standard)

	return
}

func (c *CardItem) fillClasses(db *sql.DB) (err error) {
	query := `
		SELECT name FROM classes
		INNER JOIN class ON classes.class = class.class
		WHERE classes.card = ?;
	`

	rows, err := db.Query(query, c.ID)
	if err != nil {
		return
	}
	defer rows.Close()

	names := make([]string, 0)
	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			return
		}
		names = append(names, name)
	}

	if err = rows.Err(); err != nil {
		return
	}

	c.Classes = names

	return
}

func (c *CardItem) fillKeywords(db *sql.DB) (err error) {
	query := `
		SELECT name, ref FROM mechanism
		INNER JOIN keyword ON mechanism.keyword = keyword.keyword
		WHERE mechanism.card = ?;
	`

	rows, err := db.Query(query, c.ID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var ref string
		if err = rows.Scan(&name, &ref); err != nil {
			return
		}
		item := KeywordItem{
			Name: name,
			Text: ref,
		}
		c.Keywords = append(c.Keywords, item)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}

func (c *CardItem) fillParents(db *sql.DB) (err error) {
	query := `
		SELECT parent FROM family
		WHERE child = ?
	`

	rows, err := db.Query(query, c.ID)
	if err != nil {
		return
	}
	defer rows.Close()

	ids := make([]int, 0)
	for rows.Next() {
		var id int
		if err = rows.Scan(&id); err != nil {
			return
		}
		ids = append(ids, id)
	}

	if err = rows.Err(); err != nil {
		return
	}

	for _, id := range ids {
		item, _ := newInnerCardItem(db, id)
		c.Parents = append(c.Parents, item)
	}

	return
}

func (c *CardItem) fillChildren(db *sql.DB) (err error) {
	query := `
		SELECT child FROM family
		WHERE parent = ?
	`

	rows, err := db.Query(query, c.ID)
	if err != nil {
		return
	}
	defer rows.Close()

	ids := make([]int, 0)
	for rows.Next() {
		var id int
		if err = rows.Scan(&id); err != nil {
			return
		}
		ids = append(ids, id)
	}

	if err = rows.Err(); err != nil {
		return
	}

	for _, id := range ids {
		item, _ := newInnerCardItem(db, id)
		c.Children = append(c.Children, item)
	}

	return
}
