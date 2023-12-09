package day9

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func GetSumOfPastValues(documentName string) (sum int) {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		steps := parseSteps(scanner.Text())
		diffs := map[int][]int{0: steps}
		count := 0
		for !allZeros(steps) {
			steps = calculateNextSteps(steps)
			count += 1
			diffs[count] = steps
		}
		slices.Reverse(diffs[count])
		diffs[count] = append(diffs[count], 0)
		curr := 0
		for i := count - 1; i >= 0; i-- {
			slices.Reverse(diffs[i])
			next := diffs[i][len(diffs[i])-1] - curr
			diffs[i] = append(diffs[i], next)
			curr = next
		}
		sum += curr
		fmt.Println(diffs)
	}
	return
}

func GetSumOfFutureValues(documentName string) (sum int) {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		steps := parseSteps(scanner.Text())
		diffs := map[int][]int{0: steps}
		count := 0
		for !allZeros(steps) {
			steps = calculateNextSteps(steps)
			count += 1
			diffs[count] = steps
		}
		diffs[count] = append(diffs[count], 0)
		curr := 0
		for i := count - 1; i >= 0; i-- {
			next := diffs[i][len(diffs[i])-1] + curr
			diffs[i] = append(diffs[i], next)
			curr = next
		}
		sum += curr
	}
	return
}

func allZeros(steps []int) bool {
	for _, step := range steps {
		if step != 0 {
			return false
		}
	}
	return true
}

func calculateNextSteps(steps []int) (differences []int) {
	for i, step := range steps {
		if i+1 < len(steps) {
			differences = append(differences, steps[i+1]-step)
		}
	}
	return
}

func parseSteps(line string) []int {
	return atoi(strings.Fields(line))
}

func atoi(strings []string) (ints []int) {
	for _, str := range strings {
		num, _ := strconv.Atoi(str)
		ints = append(ints, num)
	}
	return
}
