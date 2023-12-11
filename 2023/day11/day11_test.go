package day11

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestGetSumOfGalaxyDistancesTwo(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetSumOfGalaxyDistances("test_data/"+testFile.Name(), 2)
		t.Logf(strconv.Itoa(actual))
	}
}

func TestGetSumOfGalaxyDistancesTen(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetSumOfGalaxyDistances("test_data/"+testFile.Name(), 10)
		t.Logf(strconv.Itoa(actual))
	}
}

func TestGetSumOfGalaxyDistancesHundred(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetSumOfGalaxyDistances("test_data/"+testFile.Name(), 100)
		t.Logf(strconv.Itoa(actual))
	}
}

func TestGetSumOfGalaxyDistancesMillion(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetSumOfGalaxyDistances("test_data/"+testFile.Name(), 1000000)
		t.Logf(strconv.Itoa(actual))
	}
}
