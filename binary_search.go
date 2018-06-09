package main

import (
	"os"
	"log"
	"io"
	"fmt"
	"bytes"
	"strings"
	"strconv"
	"unicode"
	"unicode/utf8"
)

type CompareMode int

const (
	IgnoreWhitespace CompareMode = 1 << iota
	CaseInsensitive CompareMode = 1 << iota
	Numeric CompareMode = 1 << iota
)

type BinarySearch struct {
	fileName string
	file     *os.File
	reverse bool
	compareMode CompareMode
}

func NewBinarySearch(fileName string, reverse bool, compareMode CompareMode) BinarySearch {
	return BinarySearch{fileName: fileName, reverse: reverse, compareMode: compareMode}
}

func (search *BinarySearch) checkMatch(match string, matchPos int, data []byte) bool {
	if search.compareMode & CaseInsensitive == CaseInsensitive {
		return matchPos < len(match) && bytes.ToLower(data)[0] == match[matchPos]
	} else {
		return matchPos < len(match) && data[0] == match[matchPos]
	}
}

func (search *BinarySearch) CompareBytes(match string) int {
	if search.compareMode & CaseInsensitive == CaseInsensitive {
		match = strings.ToLower(match)
	}
	if search.compareMode & IgnoreWhitespace == IgnoreWhitespace {
		match = strings.TrimLeft(match, " \t")
	}
	data := make([]byte, 1)
	var matchPos = 0
	var byteErr error = nil
	var err error = nil
	if search.compareMode & IgnoreWhitespace == IgnoreWhitespace {
		for _, err = search.file.Read(data); data[0] != '\n' && (data[0] == ' ' || data[0] == '\t'); _, err = search.file.Read(data) {
			byteErr = err
		}
	} else {
		_, err = search.file.Read(data)
	}
	for err = nil; search.checkMatch(match, matchPos, data); _, err = search.file.Read(data) {
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

	if search.compareMode & CaseInsensitive == CaseInsensitive {
		return int(match[matchPos]) - int(bytes.ToLower(data)[0])
	}
	return int(match[matchPos]) - int(data[0])
}

func (search *BinarySearch) CompareNumeric(match string) int {
	matchVal, err := strconv.ParseFloat(match, 64)
	var buffer bytes.Buffer
	data := make([]byte, 1)
	for _, err = search.file.Read(data); err != io.EOF && data[0] != '\n'; _, err = search.file.Read(data) {
		if buffer.Len() == 0 && (data[0] == ' ' || data[0] == '\t') {
			continue
		}
		r, _ := utf8.DecodeRune(data)

		if data[0] != '-' && data[0] != '.' && !unicode.IsDigit(r) {
			break
		}
		buffer.WriteByte(data[0])
	}
	if buffer.Len() == 0 {
		return -1
	}

	bufferString := buffer.String()
	var bufferVal, _ = strconv.ParseFloat(bufferString, 64)

	if matchVal < bufferVal {
		return -1
	}
	if matchVal > bufferVal {
		return 1
	}
	return 0
}

func (search *BinarySearch) Compare(match string) int {
	var cmp int

	if search.compareMode & Numeric == Numeric {
		cmp = search.CompareNumeric(match)
	} else {
		cmp = search.CompareBytes(match)
	}
	if search.reverse {
		cmp = -cmp
	}
	return cmp
}

func (search *BinarySearch) PrintMatch(start int64, match string) {
	if start >= 0 {
		search.file.Seek(start, os.SEEK_SET)
	}
	data := make([]byte, 1)
	var full_data []byte

	for count := 0; start >= 0 && search.Compare(match) == 0; count++ {
		search.file.Seek(start, os.SEEK_SET)
		var isEof = false
		for _, err := search.file.Read(data); !isEof; _, err = search.file.Read(data) {
			isEof = err == io.EOF
			if isEof {
				break
			}
			start++
			full_data = append(full_data, data[0])
			if data[0] == '\n' {
				break
			}
		}
		if isEof {
			break
		}
	}
	fmt.Printf("%s", string(full_data))

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
