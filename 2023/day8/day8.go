package day8

import (
	"bufio"
	"math"
	"os"
	"strings"
)

func GetRequiredStepsFromAllANodes(documentName string) (steps int) {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := parseInstructions(scanner.Text())
	scanner.Scan()
	network := map[string][]string{}
	var currNodes []string
	for scanner.Scan() {
		node, nextElements := parseNode(scanner.Text())
		if node[2] == 'A' {
			currNodes = append(currNodes, node)
		}
		network[node] = nextElements
	}
	countsMap := map[string]int{}
	for _, currNode := range currNodes {
		currSteps := 0
		nextNode := currNode
		for i := 0; nextNode[2] != 'Z'; i = (i + 1) % len(instructions) {
			nextNode = getNextElement(nextNode, instructions[i], network)
			currSteps += 1
		}
		countsMap[currNode] = currSteps
	}
	return getCombinedLcm(getValues(countsMap))
}

func getCombinedLcm(nums []int) int {
	lcm := 1
	for _, num := range nums {
		lcm = getLcm(lcm, num)
	}
	return lcm
}

func getLcm(a, b int) (lcm int) {
	return int(math.Abs(float64(a*b)) / float64(getGcd(a, b)))
}

func getGcd(a, b int) int {
	if b == 0 {
		return a
	}
	return getGcd(b, a%b)
}

func getValues(m map[string]int) (values []int) {
	for _, value := range m {
		values = append(values, value)
	}
	return
}

func GetRequiredStepsFromAAA(documentName string) (steps int) {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := parseInstructions(scanner.Text())
	scanner.Scan()
	network := map[string][]string{}
	for scanner.Scan() {
		node, nextElements := parseNode(scanner.Text())
		network[node] = nextElements
	}
	currNode := "AAA"
	for i := 0; currNode != "ZZZ"; i = (i + 1) % len(instructions) {
		currNode = getNextElement(currNode, instructions[i], network)
		steps += 1
	}
	return
}

func parseInstructions(line string) (instructions []int) {
	for _, char := range line {
		if char == 'L' {
			instructions = append(instructions, 0)
		} else {
			instructions = append(instructions, 1)
		}
	}
	return
}

func parseNode(line string) (string, []string) {
	node := strings.Split(line, " = ")[0]
	nextElements := strings.Split(line, " = ")[1]
	leftElement := strings.Split(nextElements[1:len(nextElements)-1], ", ")[0]
	rightElement := strings.Split(nextElements[1:len(nextElements)-1], ", ")[1]
	return node, []string{leftElement, rightElement}
}

func getNextElement(currNode string, nextInstruction int, network map[string][]string) string {
	return network[currNode][nextInstruction]
}
