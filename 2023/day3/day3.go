package day3

import (
	"bufio"
	"os"
	"strconv"
)

func GetGearRatiosSum(documentName string) int {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	var schematic []string
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, line)
	}
	return parseSchematicForGearRatios(schematic)
}

func parseSchematicForGearRatios(schematic []string) int {
	gearCoordsMap := make(map[[2]int][]int)
	for i, line := range schematic {
		currNum := 0
		numStart := -1
		for j, char := range line {
			if digit, err := strconv.Atoi(string(char)); err == nil {
				if currNum == 0 {
					numStart = j
				}
				currNum = currNum*10 + digit
			} else if currNum != 0 {
				if hasGear, x, y := isGearRatio(i, j, numStart, schematic); hasGear {
					currList, coordsExist := gearCoordsMap[[2]int{x, y}]
					if coordsExist {
						gearCoordsMap[[2]int{x, y}] = append(currList, currNum)
					} else {
						gearCoordsMap[[2]int{x, y}] = []int{currNum}
					}
				}
				currNum = 0
				numStart = -1
			}
		}
		if hasGear, x, y := isGearRatio(i, len(schematic[i]), numStart, schematic); hasGear && numStart >= 0 {
			currList, coordsExist := gearCoordsMap[[2]int{x, y}]
			if coordsExist {
				gearCoordsMap[[2]int{x, y}] = append(currList, currNum)
			} else {
				gearCoordsMap[[2]int{x, y}] = []int{currNum}
			}
		}
	}
	return collectValidGearRatios(gearCoordsMap)
}

func collectValidGearRatios(gearCoordsMap map[[2]int][]int) (sum int) {
	for _, partNumbers := range gearCoordsMap {
		if len(partNumbers) == 2 {
			sum += product(partNumbers)
		}
	}
	return
}

func product(values []int) int {
	res := 1
	for _, value := range values {
		res *= value
	}
	return res
}

func isGearRatio(i, j, start int, schematic []string) (bool, int, int) {
	if start > 0 && isGear(schematic[i][start-1]) {
		return true, i, start - 1
	} else if j < len(schematic[i])-1 && isGear(schematic[i][j]) {
		return true, i, j
	}
	for _, k := range makeRange(max(0, start-1), min(len(schematic[i])-1, j)) {
		if i > 0 && isGear(schematic[i-1][k]) {
			return true, i - 1, k
		} else if i < len(schematic)-1 && isGear(schematic[i+1][k]) {
			return true, i + 1, k
		}
	}
	return false, i, j
}

func isGear(char uint8) bool {
	return char == '*'
}

func GetPartNumbersSum(documentName string) int {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	var schematic []string
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, line)
	}
	return parseSchematicForPartNumbers(schematic)
}

func parseSchematicForPartNumbers(schematic []string) int {
	sum := 0
	for i, line := range schematic {
		currNum := 0
		numStart := -1
		for j, char := range line {
			if digit, err := strconv.Atoi(string(char)); err == nil {
				if currNum == 0 {
					numStart = j
				}
				currNum = currNum*10 + digit
			} else if currNum != 0 {
				if isPartNumber(i, j, numStart, schematic) {
					sum += currNum
				}
				currNum = 0
				numStart = -1
			}
		}
		if numStart >= 0 && isPartNumber(i, len(schematic[i]), numStart, schematic) {
			sum += currNum
		}
	}
	return sum
}

func isPartNumber(i, j, start int, schematic []string) bool {
	if (start > 0 && isSymbol(schematic[i][start-1])) || (j < len(schematic[i])-1 && isSymbol(schematic[i][j])) {
		return true
	}
	for _, k := range makeRange(max(0, start-1), min(len(schematic[i])-1, j)) {
		if (i > 0 && isSymbol(schematic[i-1][k])) || (i < len(schematic)-1 && isSymbol(schematic[i+1][k])) {
			return true
		}
	}
	return false
}

func isSymbol(char uint8) bool {
	if _, err := strconv.Atoi(string(char)); err != nil {
		return char != '.'
	}
	return false
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
