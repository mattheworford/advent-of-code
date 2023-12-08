package day6

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetWinningVariations(documentName string) int {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	times := atoi(strings.Fields(strings.Split(scanner.Text(), ": ")[1]))
	scanner.Scan()
	records := atoi(strings.Fields(strings.Split(scanner.Text(), ": ")[1]))
	product := 1
	for i, time := range times {
		numWinningVariations := 0
		for j := 1; j < time; j++ {
			distance := getDistance(j, time)
			if distance > records[i] {
				numWinningVariations += 1
			}
		}
		product *= numWinningVariations
	}
	return product
}

func getDistance(holdingTime, totalTime int) int {
	remainingTime := totalTime - holdingTime
	return remainingTime * holdingTime
}

func atoi(strings []string) (ints []int) {
	for _, str := range strings {
		num, _ := strconv.Atoi(str)
		ints = append(ints, num)
	}
	return
}
