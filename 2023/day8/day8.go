package day8

import (
	"bufio"
	"os"
	"strings"
)

func GetRequiredSteps(documentName string) (steps int) {
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
	for i := 0; currNode != "ZZZ"; i++ {
		i = i % len(instructions)
		currNode = network[currNode][instructions[i]]
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
