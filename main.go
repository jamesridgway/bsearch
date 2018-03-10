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
	app.Version = "0.0.1"
	app.ArgsUsage = "SEARCH_KEY FILENAME"

	app.Action = func(c *cli.Context) error {

		if c.NArg() != 2 {
			fmt.Println("Usage: bsearch [global options] command [command options] SEARCH_KEY FILENAME")
			fmt.Println("Try 'bsearch --help' for more information.")
			os.Exit(1)
		}

		searchCriteria := c.Args().Get(0)
		fileName := c.Args().Get(1)

		bsearch := NewBinarySearch(fileName)

		startPosition := bsearch.FindStart(searchCriteria)

		if startPosition > -1 {
			bsearch.PrintMatch(startPosition, searchCriteria)
		}

		return nil
	}

	app.Run(os.Args)

}
