package main

import (
	"cardcaptor/cmd"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	subCmd, err := cmd.Parse(os.Args)
	if err != nil {
		cmd.PrintUsage()
		os.Exit(0)
	}

	if err = subCmd.ParseFlags(os.Args[2:]); err != nil {
		os.Exit(0)
	}

	if err = subCmd.Run(); err != nil {
		log.Fatal(err)
	}
}
