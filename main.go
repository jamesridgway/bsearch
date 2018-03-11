package main

import (
	"github.com/urfave/cli"
	"os"
	"fmt"
)

func main() {

	app := cli.NewApp()
	app.Name = "bsearch"
	app.Usage = "utility for binary searching a sorted file for lines that start with the search key"
	app.Version = "1.0.1"
	app.ArgsUsage = "SEARCH_KEY FILENAME"

	var reverse = false
	var ignoreWhitespace = false
	var caseInsensitive = false
	var numeric = false
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "r,reverse",
			Usage:       "the reverse flag indicates the file is sorted in descending order",
			Destination: &reverse,
		},
		cli.BoolFlag{
			Name:        "i,ignore-case",
			Usage:       "case insensitive",
			Destination: &caseInsensitive,
		},
		cli.BoolFlag{
			Name:        "t,trim",
			Usage:       "ignore whitespace",
			Destination: &ignoreWhitespace,
		},
		cli.BoolFlag{
			Name:        "n,numeric",
			Usage:       "use numeric comparison",
			Destination: &numeric,
		},
	}

	app.Action = func(c *cli.Context) error {

		if c.NArg() != 2 {
			fmt.Println("Usage: bsearch [options] SEARCH_KEY FILENAME")
			fmt.Println("Try 'bsearch --help' for more information.")
			os.Exit(1)
		}

		searchCriteria := c.Args().Get(0)
		fileName := c.Args().Get(1)

		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			fmt.Println("bsearch: no such file")
			os.Exit(1)
		}

		var compareMode CompareMode = 0
		if ignoreWhitespace {
			compareMode = compareMode | IgnoreWhitespace
		}
		if caseInsensitive {
			compareMode = compareMode | CaseInsensitive
		} else if numeric {
			compareMode = compareMode | Numeric
		}

		bsearch := NewBinarySearch(fileName, reverse, compareMode)

		startPosition := bsearch.FindStart(searchCriteria)

		if startPosition > -1 {
			bsearch.PrintMatch(startPosition, searchCriteria)
		}

		return nil
	}

	app.Run(os.Args)

}
