package cmd

import (
	"cardcaptor/lib"
	"flag"
	"log"
)

type storeCmd struct {
	akid      string
	secretKey string
	dbPath    string
	region    string
	tableName string
	fs        *flag.FlagSet
}

func (s *storeCmd) ParseFlags(args []string) error {
	s.fs = flag.NewFlagSet("store", flag.ContinueOnError)
	s.fs.StringVar(&s.secretKey, "secret", "", "(Required) AWS IAM user SecretKey")
	s.fs.StringVar(&s.dbPath, "db", "./hearthstone.db", "Crawled sqlite db file path")
	s.fs.StringVar(&s.region, "region", "ap-northeast-2", "AWS region.")
	s.fs.StringVar(&s.akid, "akid", "", "(Required) AWS IAM user AccessKey")
	s.fs.StringVar(&s.tableName, "table", "hearthstone", "Table name to created")

	if err := s.fs.Parse(args); err != nil {
		return err
	}

	if err := isFlagsPassed(s.fs, "akid", "secret"); err != nil {
		s.PrintUsage()
		return err
	}

	return nil
}

func (s *storeCmd) PrintUsage() {
	s.fs.PrintDefaults()
}

func (s *storeCmd) Run() error {
	ddb := lib.NewDDB(s.region, s.akid, s.secretKey)
	db := lib.NewDB(s.dbPath)
	count, _ := lib.CardsCount(db)

	for i := 0; true; i++ {

		id, err := lib.CardIDByOffset(db, i)
		if err != nil {
			break
		}
		item, err := lib.NewCardItem(db, id)
		if err != nil {
			return err
		}
		err = item.PutItem(ddb, s.tableName)
		if err != nil {
			return err
		}

		log.Printf("%d/%d: %s Inserted successfully\n", i+1, count, item.Name)
	}

	return nil
}
