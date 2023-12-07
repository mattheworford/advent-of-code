package day2

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetValidGameRecords(documentName string) int {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if isValidGameInfo(line) {
			gameInfoAndRecord := strings.Split(line, ": ")
			gameId, err := strconv.Atoi(strings.Fields(gameInfoAndRecord[0])[1])
			if err == nil {
				record := gameInfoAndRecord[1]
				if isValidGame(record) {
					sum += gameId
				}
			}
		}
	}
	return sum
}

func isValidGame(game string) bool {
	sets := strings.Split(game, "; ")
	for _, set := range sets {
		if !isValidSet(set) {
			return false
		}
	}
	return true
}

func isValidSet(set string) bool {
	cubeCounts := strings.Split(set, ", ")
	for _, cubeCount := range cubeCounts {
		if !isValidCubeCount(cubeCount) {
			return false
		}
	}
	return true
}

func isValidCubeCount(cubeCount string) bool {
	count, err := strconv.Atoi(strings.Fields(cubeCount)[0])
	if err != nil {
		return false
	}
	color := strings.Fields(cubeCount)[1]
	switch color {
	case "red":
		return count <= 12
	case "green":
		return count <= 13
	case "blue":
		return count <= 14
	default:
		return false
	}
}

func isValidGameInfo(line string) bool {
	return regexp.MustCompile(`^Game \d*:(( \d* (red|blue|green),?)*;?)*$`).MatchString(line)
}
