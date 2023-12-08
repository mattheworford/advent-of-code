package day6

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetWinningVariationsFromSingleTime(documentName string) int {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	splitTime := strings.Fields(strings.Split(scanner.Text(), ": ")[1])
	time, _ := strconv.Atoi(strings.Join(splitTime, ""))
	scanner.Scan()
	splitRecord := strings.Fields(strings.Split(scanner.Text(), ": ")[1])
	record, _ := strconv.Atoi(strings.Join(splitRecord, ""))
	return getNumWinningVariations(time, record)
}

func GetWinningVariationsFromMultipleTimes(documentName string) int {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	times := atoi(strings.Fields(strings.Split(scanner.Text(), ": ")[1]))
	scanner.Scan()
	records := atoi(strings.Fields(strings.Split(scanner.Text(), ": ")[1]))
	product := 1
	for i, time := range times {
		product *= getNumWinningVariations(time, records[i])
	}
	return product
}

func getNumWinningVariations(time, record int) (res int) {
	for i := 1; i < time; i++ {
		distance := getDistance(i, time)
		if distance > record {
			res = time - 1 - ((i - 1) * 2)
			break
		}
	}
	return
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
