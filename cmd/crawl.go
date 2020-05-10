package cmd

import (
	"cardcaptor/lib"
	"database/sql"
	"flag"
	"log"
	"time"
)

type crawlCmd struct {
	accessToken string
	dbPath      string
	fs          *flag.FlagSet
}

func (c *crawlCmd) ParseFlags(args []string) error {

	c.fs = flag.NewFlagSet("crawl", flag.ContinueOnError)
	c.fs.StringVar(&c.dbPath, "db", "./hearthstone.db", "path for newly created db, includes extension")
	c.fs.StringVar(&c.accessToken, "key", "", "(Required) AccessKey which is provided From Blizzard API")

	if err := c.fs.Parse(args); err != nil {
		return err
	}

	if err := isFlagsPassed(c.fs, "key"); err != nil {
		c.PrintUsage()
		return err
	}

	return nil
}

func (c *crawlCmd) PrintUsage() {
	c.fs.PrintDefaults()
}

func (c *crawlCmd) Run() error {
	meta, err := lib.CrawlMetadata(c.accessToken)
	if err != nil {
		return err
	}

	lib.CreateNewDB(c.dbPath)
	db, _ := sql.Open("sqlite3", c.dbPath)
	defer db.Close()

	err = meta.Insert(db)
	if err != nil {
		return err
	}

	for _, id := range meta.ClassCards() {
		card, err := lib.FetchByID(id, c.accessToken)
		err = card.Insert(db)
		if err != nil {
			return err
		}
	}

	cardResp, err := lib.RequestCards(1, c.accessToken)
	if err != nil {
		return err
	}

	for {
		log.Printf("current page: %d/%d", cardResp.Page, cardResp.PageCount)
		for _, card := range cardResp.Cards {
			err = card.Insert(db)
			if err != nil {
				return err
			}
		}
		if !cardResp.HasNext() {
			break
		}
		time.Sleep(time.Second * 5)
		cardResp, err = cardResp.Next(c.accessToken)
		if err != nil {
			return err
		}
	}

	err = meta.InsertArena(db)
	if err != nil {
		return err
	}

	lib.InsertMissingData(db)
	lib.VacuumDB(db)
	return nil
}
