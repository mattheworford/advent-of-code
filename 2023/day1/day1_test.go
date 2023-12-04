package day1

import (
	"strconv"
	"testing"
)

var testFiles = []string{
	"test1.txt",
	"test2.txt",
	"test3.txt",
}

func TestGetTrebuchetCalibrationValues(t *testing.T) {
	for _, testFile := range testFiles {
		actual := GetTrebuchetCalibrationValues(testFile)
		t.Logf(strconv.Itoa(actual))
	}
}
