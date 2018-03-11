package main

import "testing"

func TestFindStartNoMatch(t *testing.T) {
	bsearch := NewBinarySearch("test/data/sorted_udids.txt", false, 0)
	var result = bsearch.FindStart("not-a-udid")
	if result != -1 {
		t.Errorf("result should be -1 for a non-match, got: %d", result)
	}
}

func TestFindStartEmptyFile(t *testing.T) {
	bsearch := NewBinarySearch("test/data/empty_file.txt", false, 0)
	var result = bsearch.FindStart("hello-world")
	if result != -1 {
		t.Errorf("result should be -1 for an empty file, got: %d", result)
	}
}

func TestFindStartFindFirst(t *testing.T) {
	bsearch := NewBinarySearch("test/data/sorted_udids.txt", false, 0)
	var result = bsearch.FindStart("00e8330b-f67c-4f73-bbf6-06554816048f")
	if result != 0 {
		t.Errorf("result should first hash in the data file, expecting 0, got: %d", result)
	}
}

func TestFindStartFindLast(t *testing.T) {
	bsearch := NewBinarySearch("test/data/sorted_udids.txt", false, 0)
	var result = bsearch.FindStart("ff8081b8-ca20-40ce-8675-426e442a0f8e")
	if result != 3663 {
		t.Errorf("result should first last in the data file, expecting 0, got: %d", result)
	}
}

func TestFindStart(t *testing.T) {
	bsearch := NewBinarySearch("test/data/sorted_udids.txt", false, 0)
	var result = bsearch.FindStart("9a6c81f7-eb5d-4c6a-a840-6e0e12dbe003")
	if result != 2516 {
		t.Errorf("expecting '9a6c81f7-eb5d-4c6a-a840-6e0e12dbe003' to have starting position 2516, got: %d",
			result)
	}

	var resultCaseMismatch = bsearch.FindStart("9A6C81f7-EB5D-4C6A-A840-6E0E12DBE003")
	if resultCaseMismatch != -1 {
		t.Errorf("expecting '9A6C81f7-EB5D-4C6A-A840-6E0E12DBE003' to have starting position -1, got: %d",
			result)
	}

	bsearch = NewBinarySearch("test/data/sorted_udids.txt", false, IgnoreWhitespace)
	var resultIgnoreLeadingWhitespace = bsearch.FindStart(" \t9a6c81f7-eb5d-4c6a-a840-6e0e12dbe003")
	if resultIgnoreLeadingWhitespace != 2516 {
		t.Errorf("expecting ' 9a6c81f7-eb5d-4c6a-a840-6e0e12dbe003' to have starting position 2516, got: %d",
			result)
	}

}

func TestFindStartCaseInsensitive(t *testing.T) {
	bsearch := NewBinarySearch("test/data/sorted_udids.txt", false, CaseInsensitive)
	var result = bsearch.FindStart("9a6c81f7-eb5d-4c6a-a840-6e0e12dbe003")
	if result != 2516 {
		t.Errorf("expecting '9a6c81f7-eb5d-4c6a-a840-6e0e12dbe003' to have starting position 2516, got: %d",
			result)
	}

	var resultCaseMismatch = bsearch.FindStart("9A6C81f7-EB5D-4C6A-A840-6E0E12DBE003")
	if resultCaseMismatch != 2516 {
		t.Errorf("expecting '9A6C81f7-EB5D-4C6A-A840-6E0E12DBE003' to have starting position 2516, got: %d",
			result)
	}
}

func TestFindStartWithReverse(t *testing.T) {
	bsearch := NewBinarySearch("test/data/reverse_sorted_udids.txt", true, 0)
	var result = bsearch.FindStart("9a6c81f7-eb5d-4c6a-a840-6e0e12dbe003")
	if result != 1147 {
		t.Errorf("expecting '9a6c81f7-eb5d-4c6a-a840-6e0e12dbe003' to have starting position 1147, got: %d",
			result)
	}
}

func TestFindStartSingleCharacter(t *testing.T) {
	bsearch := NewBinarySearch("test/data/sorted_alphabets.txt", false, 0)
	var result = bsearch.FindStart("k")
	if result != 8 {
		t.Errorf("expecting 'k' to have starting position 8, got: %d",
			result)
	}
}

func TestFindStartSingleCharacterWhitespace(t *testing.T) {
	bsearch := NewBinarySearch("test/data/sorted_alphabets_whitespace.txt", false, IgnoreWhitespace)
	var result = bsearch.FindStart(" \tk")
	if result != 12 {
		t.Errorf("expecting 'k' to have starting position 12, got: %d",
			result)
	}
}

func TestFindStartSingleCharacterBeyondEndOfFile(t *testing.T) {
	bsearch := NewBinarySearch("test/data/sorted_alphabets.txt", false, 0)
	var result = bsearch.FindStart("z")
	if result != -1 {
		t.Errorf("expecting 'z' to have starting position -1, got: %d",
			result)
	}
}

func TestFindStartNumeric(t *testing.T) {
	bsearch := NewBinarySearch("test/data/sorted_alphabets_numbered.txt", false, Numeric)
	var result= bsearch.FindStart("112")
	if result != 37 {
		t.Errorf("expecting '112' to have starting position 37, got: %d",
			result)
	}
}