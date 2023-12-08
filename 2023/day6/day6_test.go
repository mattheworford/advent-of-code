package day6

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
		actual := GetWinningVariations("test_data/" + testFile.Name())
		t.Logf(strconv.Itoa(actual))
	}
}
