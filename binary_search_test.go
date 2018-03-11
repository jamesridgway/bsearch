package main

import "testing"

func TestFindStartNoMatch(t *testing.T) {
	bsearch := NewBinarySearch("test/data/sorted_udids.txt", false)
	var result = bsearch.FindStart("not-a-udid")
	if result != -1 {
		t.Errorf("result should be -1 for a non-match, got: %d", result)
	}
}

func TestFindStartEmptyFile(t *testing.T) {
	bsearch := NewBinarySearch("test/data/empty_file.txt", false)
	var result = bsearch.FindStart("hello-world")
	if result != -1 {
		t.Errorf("result should be -1 for an empty file, got: %d", result)
	}
}

func TestFindStartFindFirst(t *testing.T) {
	bsearch := NewBinarySearch("test/data/sorted_udids.txt", false)
	var result = bsearch.FindStart("00e8330b-f67c-4f73-bbf6-06554816048f")
	if result != 0 {
		t.Errorf("result should first hash in the data file, expecting 0, got: %d", result)
	}
}

func TestFindStartFindLast(t *testing.T) {
	bsearch := NewBinarySearch("test/data/sorted_udids.txt", false)
	var result = bsearch.FindStart("ff8081b8-ca20-40ce-8675-426e442a0f8e")
	if result != 3663 {
		t.Errorf("result should first last in the data file, expecting 0, got: %d", result)
	}
}

func TestFindStart(t *testing.T) {
	bsearch := NewBinarySearch("test/data/sorted_udids.txt", false)
	var result = bsearch.FindStart("9a6c81f7-eb5d-4c6a-a840-6e0e12dbe003")
	if result != 2516 {
		t.Errorf("expecting '9a6c81f7-eb5d-4c6a-a840-6e0e12dbe003' to have starting position 2516, got: %d",
			result)
	}
}

func TestFindStartWithReverse(t *testing.T) {
	bsearch := NewBinarySearch("test/data/reverse_sorted_udids.txt", true)
	var result = bsearch.FindStart("9a6c81f7-eb5d-4c6a-a840-6e0e12dbe003")
	if result != 1147 {
		t.Errorf("expecting '9a6c81f7-eb5d-4c6a-a840-6e0e12dbe003' to have starting position 1147, got: %d",
			result)
	}
}

func TestFindStartSingleCharacter(t *testing.T) {
	bsearch := NewBinarySearch("test/data/sorted_alphabets.txt", false)
	var result = bsearch.FindStart("k")
	if result != 8 {
		t.Errorf("expecting 'k' to have starting position 8, got: %d",
			result)
	}
}

func TestFindStartSingleCharacterBeyondEndOfFile(t *testing.T) {
	bsearch := NewBinarySearch("test/data/sorted_alphabets.txt", false)
	var result = bsearch.FindStart("z")
	if result != -1 {
		t.Errorf("expecting 'z' to have starting position -1, got: %d",
			result)
	}
}
