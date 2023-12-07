package day2

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestGetValidGameRecords(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetValidGameRecords("test_data/" + testFile.Name())
		t.Logf(strconv.Itoa(actual))
	}
}

func TestGetGamePowers(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetGamePowers("test_data/" + testFile.Name())
		t.Logf(strconv.Itoa(actual))
	}
}
