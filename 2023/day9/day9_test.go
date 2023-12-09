package day9

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestGetSumOfFutureValues(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetSumOfFutureValues("test_data/" + testFile.Name())
		t.Logf(strconv.Itoa(actual))
	}
}

func TestGetSumOfPastValues(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetSumOfPastValues("test_data/" + testFile.Name())
		t.Logf(strconv.Itoa(actual))
	}
}
