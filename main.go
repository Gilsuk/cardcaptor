package main

import (
	"cardcaptor/lib"
	"database/sql"
	"flag"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbPath := flag.String("db", "./hearthstone.db", "path for newly created db, includes extension")
	accessToken := flag.String("key", "", "AccessKey which is provided From Blizzard API (Required)")
	flag.Parse()
	if !isFlagPassed("key") {
		flag.Usage()
		return
	}

	meta, err := lib.CrawlMetadata(*accessToken)
	if err != nil {
		log.Fatal(err)
	}

	lib.CreateNewDB(*dbPath)
	db, _ := sql.Open("sqlite3", *dbPath)
	defer db.Close()

	err = meta.Insert(db)
	if err != nil {
		log.Fatal(err)
	}

	cardResp, err := lib.RequestCards(1, *accessToken)
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Printf("current page: %d/%d", cardResp.Page, cardResp.PageCount)
		for _, card := range cardResp.Cards {
			err = card.Insert(db)
			if err != nil {
				log.Fatal(err)
			}
		}
		if !cardResp.HasNext() {
			break
		}
		time.Sleep(time.Second * 5)
		cardResp, err = cardResp.Next(*accessToken)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = meta.InsertArena(db)
	if err != nil {
		log.Fatal(err)
	}

	lib.VacuumDB(db)
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
