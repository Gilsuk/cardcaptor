package cmd

import (
	"cardcaptor/lib"
	"flag"
)

type structCmd struct {
	akid      string
	secretKey string
	tableName string
	region    string
	rcu       int
	wcu       int
	fs        *flag.FlagSet
}

func (s *structCmd) ParseFlags(args []string) error {
	s.fs = flag.NewFlagSet("struct", flag.ContinueOnError)
	s.fs.StringVar(&s.secretKey, "secret", "", "(Required) AWS IAM user SecretKey")
	s.fs.StringVar(&s.tableName, "table", "hearthstone", "Table name to created")
	s.fs.StringVar(&s.region, "region", "ap-northeast-2", "AWS region")
	s.fs.StringVar(&s.akid, "akid", "", "(Required) AWS IAM user AccessKey")
	s.fs.IntVar(&s.rcu, "rcu", 10, "Provisioned read capacity units(RCU)")
	s.fs.IntVar(&s.wcu, "wcu", 10, "Provisioned write capacity units(WCU)")

	if err := s.fs.Parse(args); err != nil {
		return err
	}

	if err := isFlagsPassed(s.fs, "akid", "secret"); err != nil {
		s.PrintUsage()
		return err
	}

	return nil
}

func (s *structCmd) PrintUsage() {
	s.fs.PrintDefaults()
}

func (s *structCmd) Run() error {
	db := lib.NewDDB(s.region, s.akid, s.secretKey)
	err := lib.CreateDDBTable(db, s.tableName, int64(s.rcu), int64(s.wcu))
	if err != nil {
		return err
	}
	return nil
}
