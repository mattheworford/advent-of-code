package day4

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestGetTotalScratchcardCount(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetTotalScratchcardCount("test_data/" + testFile.Name())
		t.Logf(strconv.Itoa(actual))
	}
}

func TestGetGearRatiosSum(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetCardPointValues("test_data/" + testFile.Name())
		t.Logf(strconv.Itoa(actual))
	}
}
