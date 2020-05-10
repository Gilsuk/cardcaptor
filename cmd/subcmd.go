package cmd

import (
	"errors"
	"flag"
	"fmt"
)

// SubCommand is
type SubCommand interface {
	ParseFlags([]string) error
	PrintUsage()
	Run() error
}

func newSubCommand(cmdName string) (SubCommand, error) {
	switch cmdName {
	case "crawl":
		return &crawlCmd{}, nil
	case "store":
		return &storeCmd{}, nil
	case "struct":
		return &structCmd{}, nil
	default:
		return nil, errors.New("SubCommand not found")
	}
}

// Parse is
func Parse(args []string) (SubCommand, error) {
	if len(args) < 2 {
		return nil, errors.New("SubCommand not passed")
	}

	subCmd, err := newSubCommand(args[1])
	if err != nil {
		return nil, err
	}

	return subCmd, nil
}

// PrintUsage is
func PrintUsage() {
	subcmds := map[string]string{
		"crawl":  "Crawl data from Blizzard API server",
		"store":  "Put crawled data into DynamoDB",
		"struct": "Create DynamoDB table",
	}
	fmt.Println("--- sub commands list ---")
	for k, v := range subcmds {
		fmt.Printf("%s\t%s\n", k, v)
	}
}

func isFlagsPassed(fs *flag.FlagSet, args ...string) error {
	for _, arg := range args {
		passed := false

		fs.Visit(func(f *flag.Flag) {
			if f.Name == arg {
				passed = true
			}
		})

		if !passed {
			return fmt.Errorf("Required argument %s not passed", arg)
		}
	}
	return nil
}
