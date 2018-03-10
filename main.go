package main

import (
	"bsearch/binary_search"
	"fmt"
)

func main() {
	searchCriteria := "0866D"

	bsearch := binary_search.New("/home/james/pwned-passwords-ordered-2.0.txt")

	startPosition := bsearch.FindStart(searchCriteria)
	fmt.Println("Match starts at: ", startPosition)

	if startPosition > -1 {
		bsearch.PrintMatch(startPosition, searchCriteria)
	}
}
