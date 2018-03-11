package main

import (
	"os"
	"log"
	"io"
	"fmt"
)

type BinarySearch struct {
	fileName string
	file     *os.File
	reverse bool
}

func NewBinarySearch(fileName string, reverse bool) BinarySearch {
	return BinarySearch{fileName: fileName, reverse: reverse}
}

func (search *BinarySearch) Compare(match string) int {
	data := make([]byte, 1)
	var matchPos = 0
	var byteErr error = nil
	for _, err := search.file.Read(data); matchPos < len(match) && data[0] == match[matchPos]; _, err = search.file.Read(data) {
		byteErr = err
		matchPos++
	}

	if matchPos > len(match)-1 {
		return 0
	}
	if data[0] == '\n' {
		return 1
	}
	if byteErr == io.EOF {
		return -1
	}

	return int(match[matchPos]) - int(data[0])

}

func (search *BinarySearch) PrintMatch(start int64, match string) {
	if start >= 0 {
		search.file.Seek(start, os.SEEK_SET)
	}
	data := make([]byte, 1)
	for count := 0; start >= 0 && search.Compare(match) == 0; count++ {
		search.file.Seek(start, os.SEEK_SET)
		var isEof = false
		for _, err := search.file.Read(data); !isEof; _, err = search.file.Read(data) {
			isEof = err == io.EOF
			if isEof {
				break
			}
			start++
			fmt.Printf("%s", string(data[0]))
			if data[0] == '\n' {
				break
			}
		}
		if isEof {
			break
		}
	}

}

func (search *BinarySearch) FindStart(match string) int64 {
	var result int64 = -1

	var err error
	search.file, err = os.Open(search.fileName)
	check(err)

	stat, err := search.file.Stat()
	check(err)

	var start int64 = 0
	var end = stat.Size() - 1

	data := make([]byte, 1)

	var previous int64 = -1
	var cmp int
	for start <= end {
		var middle = (start + end) / 2
		search.file.Seek(middle, os.SEEK_SET)

		var scan = middle

		if scan != 0 {
			for _, err := search.file.Read(data); true; _, err = search.file.Read(data) {
				scan++
				if err == io.EOF || data[0] == '\n' {
					break
				}
			}
		}

		if previous != scan {
			cmp = search.Compare(match)
			if search.reverse {
				cmp = -cmp
			}
			previous = scan
		}

		if cmp < 0 {
			end = middle - 1
		} else if cmp > 0 {
			start = scan + 1
		} else {
			result = scan
			end = middle - 1
		}
	}

	return result
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
