package day1

import (
	"bufio"
	"os"
	"strconv"
)

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
	for _, char := range line {
		if digit, err := strconv.Atoi(string(char)); err == nil {
			digits = append(digits, digit)
		}
	}
	if len(digits) == 0 {
		return 0
	} else {
		return digits[0]*10 + digits[len(digits)-1]
	}
}
