package day6

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestGetWinningVariationsFromSingleTime(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetWinningVariationsFromSingleTime("test_data/" + testFile.Name())
		t.Logf(strconv.Itoa(actual))
	}
}

func TestGetWinningVariationsFromMultipleTimes(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetWinningVariationsFromMultipleTimes("test_data/" + testFile.Name())
		t.Logf(strconv.Itoa(actual))
	}
}
