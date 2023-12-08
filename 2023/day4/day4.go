package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetTotalScratchcards(documentName string) int {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	cardNumber := 1
	cardToMatches := map[int]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if containsValidCardInfo(line) {
			record := strings.Split(line, ": ")[1]
			cardToMatches[cardNumber] = getTotalMatches(record)
		}
		cardNumber += 1
	}
	fmt.Println(cardToMatches)
	return getCardsToCopies(cardToMatches)
}

func getCardsToCopies(cardToMatches map[int]int) int {
	cardToCopies := map[int]int{}
	for i := 1; i <= len(cardToMatches); i++ {
		cardToCopies = initOrAdd(cardToCopies, i, 1)
		for j := 1; j <= cardToMatches[i]; j++ {
			cardToCopies = initOrAdd(cardToCopies, i+j, cardToCopies[i])
		}
	}
	return sumValues(cardToCopies)
}

func sumValues(m map[int]int) (sum int) {
	for _, value := range m {
		sum += value
	}
	return
}

func initOrAdd(m map[int]int, key, value int) map[int]int {
	_, exists := m[key]
	if exists {
		m[key] += value
	} else {
		m[key] = value
	}
	return m
}

func getTotalMatches(card string) (copies int) {
	columns := strings.Split(card, "|")
	winningNumbers := parseNumbers(columns[0])
	actualNumbers := parseNumbers(columns[1])
	for number, _ := range actualNumbers {
		_, isWinning := winningNumbers[number]
		if isWinning {
			copies += 1
		}
	}
	return
}

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
