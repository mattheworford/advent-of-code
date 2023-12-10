package day10

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestGetStepsToFarthestPointInLoop(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetStepsToFarthestPointInLoop("test_data/" + testFile.Name())
		t.Logf(strconv.Itoa(actual))
	}
}
