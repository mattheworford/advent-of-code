package day2

import (
	"bufio"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func GetGamePowers(documentName string) int {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if isValidGameInfo(line) {
			record := strings.Split(line, ": ")[1]
			sum += getGamePower(record)
		}
	}
	return sum
}

func getGamePower(game string) int {
	splitter := func(r rune) bool {
		return strings.ContainsRune(",;", r)
	}
	cubeCounts := strings.FieldsFunc(game, splitter)
	maxCubeCounts := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, cubeCount := range cubeCounts {
		color, count := getCubeColorAndCount(cubeCount)
		if count > maxCubeCounts[color] {
			maxCubeCounts[color] = count
		}
	}
	power := 1
	for _, count := range maxCubeCounts {
		power *= count
	}
	return power
}

func getCubeColorAndCount(cubeCount string) (string, int) {
	count, err := strconv.Atoi(strings.Fields(cubeCount)[0])
	color := strings.Fields(cubeCount)[1]
	if err != nil || !slices.Contains([]string{"red", "green", "blue"}, color) {
		return "red", 0
	}
	return color, count
}

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
