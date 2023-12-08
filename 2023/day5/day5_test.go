package day5

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestGetLowestLocationOfSeeds(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetLowestLocationOfSeeds("test_data/" + testFile.Name())
		t.Logf(strconv.Itoa(actual))
	}
}

func TestGetLowestLocationOfSeedRanges(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetLowestLocationOfSeedRanges("test_data/" + testFile.Name())
		t.Logf(strconv.Itoa(actual))
	}
}
