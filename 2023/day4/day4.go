package day4

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetCardPointValues(documentName string) int {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if containsValidCardInfo(line) {
			record := strings.Split(line, ": ")[1]
			sum += getCardPointValue(record)
		}
	}
	return sum
}

func getCardPointValue(card string) (pointValue int) {
	columns := strings.Split(card, "|")
	winningNumbers := parseNumbers(columns[0])
	actualNumbers := parseNumbers(columns[1])
	for number, _ := range actualNumbers {
		_, isWinning := winningNumbers[number]
		if isWinning && pointValue == 0 {
			pointValue = 1
		} else if isWinning {
			pointValue *= 2
		}
	}
	return
}

func parseNumbers(str string) map[int]bool {
	numbers := map[int]bool{}
	for _, number := range strings.Fields(str) {
		if digit, err := strconv.Atoi(number); err == nil {
			numbers[digit] = true
		}
	}
	return numbers
}

func containsValidCardInfo(line string) bool {
	return regexp.MustCompile(`^Card *\d*:(( *\d*)*;?)* \|(( *\d*)*;?)*$`).MatchString(line)
}
