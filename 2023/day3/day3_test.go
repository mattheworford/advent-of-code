package day3

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestGetGearRatiosSum(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetGearRatiosSum("test_data/" + testFile.Name())
		t.Logf(strconv.Itoa(actual))
	}
}

func TestGetPartNumbersSum(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetPartNumbersSum("test_data/" + testFile.Name())
		t.Logf(strconv.Itoa(actual))
	}
}
