package day8

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestGetRequiredStepsFromAllANodes(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetRequiredStepsFromAllANodes("test_data/" + testFile.Name())
		t.Logf(strconv.Itoa(actual))
	}
}

func TestGetRequiredStepsFromAAA(t *testing.T) {
	testFiles, err := os.ReadDir("test_data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, testFile := range testFiles {
		actual := GetRequiredStepsFromAAA("test_data/" + testFile.Name())
		t.Logf(strconv.Itoa(actual))
	}
}
