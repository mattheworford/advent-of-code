package day1

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var digitSpellingMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func GetTrebuchetCalibrationValues(documentName string) int {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += parse(line)
	}
	return sum
}

func parse(line string) int {
	var digits []int
	currString := ""
	for _, char := range line {
		currString += string(char)
		digit := parseDigit(string(char), currString)
		if digit != 0 {
			digits = append(digits, digit)
			currString = string(char)
		}
	}
	return combineFirstAndLastDigits(digits)
}

func parseDigit(char, currString string) int {
	if digit, err := strconv.Atoi(char); err == nil {
		return digit
	} else {
		for digitSpelling := range digitSpellingMap {
			if strings.Contains(currString, digitSpelling) {
				return digitSpellingMap[digitSpelling]
			}
		}
		return 0
	}
}

func combineFirstAndLastDigits(digits []int) int {
	if len(digits) == 0 {
		return 0
	} else {
		return digits[0]*10 + digits[len(digits)-1]
	}
}
