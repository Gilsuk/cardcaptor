package lib

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"
)

// CardJSON is
type CardJSON struct {
	ID          int        `json:"id"`
	Slug        string     `json:"slug"`
	Class       string     `json:"class,omitempty"`
	Classes     []string   `json:"classes,omitempty"`
	CardType    string     `json:"type,omitempty"`
	CardSet     string     `json:"cardSet,omitempty"`
	Rarity      string     `json:"rarity,omitempty"`
	Race        string     `json:"race,omitempty"`
	Keywords    []string   `json:"keywords,omitempty"`
	Artist      string     `json:"artist,omitempty"`
	Name        string     `json:"name,omitempty"`
	Text        string     `json:"text,omitempty"`
	Flavor      string     `json:"flavor,omitempty"`
	Img         string     `json:"img,omitempty"`
	CropImg     string     `json:"cropImg,omitempty"`
	Cost        int        `json:"cost"`
	Health      int        `json:"health,omitempty"`
	Attack      int        `json:"attack,omitempty"`
	Armor       int        `json:"armor,omitempty"`
	Arena       bool       `json:"arena,omitempty"`
	Collectible bool       `json:"collectible,omitempty"`
	Standard    bool       `json:"standard,omitempty"`
	Parents     []CardJSON `json:"parents,omitempty"`
	Children    []CardJSON `json:"children,omitempty"`
}

// Export is
func (c *CardJSON) Export(dir string) {
	os.MkdirAll(dir, os.ModePerm)

	bytes, _ := json.Marshal(*c)
	ioutil.WriteFile(dir+"/"+c.Slug+".json", bytes, os.ModePerm)
}

// NewCardJSON is
func NewCardJSON(db *sql.DB, id int) (c CardJSON, err error) {
	query := `
		SELECT 
			id,slug,name,class,type,rarity,
			race,cost,health,attack,armor,
			arena,collectable,text,flavor,artist,
			cardSet,img,cropImg,standard
		FROM vNNCard
		WHERE id = ?
	`
	row := db.QueryRow(query, id)

	err = row.Scan(&c.ID, &c.Slug, &c.Name, &c.Class, &c.CardType,
		&c.Rarity, &c.Race, &c.Cost, &c.Health, &c.Attack,
		&c.Armor, &c.Arena, &c.Collectible, &c.Text, &c.Flavor,
		&c.Artist, &c.CardSet, &c.Img, &c.CropImg, &c.Standard)

	return
}
