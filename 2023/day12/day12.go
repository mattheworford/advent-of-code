package day12

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetTotalSpringArrangements(documentName string) (sum int) {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		record, groupSizes := getRecordAndGroupSizes(scanner.Text())
		sum += getNumArrangements("", record, groupSizes, 0, false)
	}
	return
}

func getRecordAndGroupSizes(line string) (record string, groupSizes []int) {
	record = strings.Fields(line)[0]
	for _, groupStr := range strings.Split(strings.Fields(line)[1], ",") {
		group, _ := strconv.Atoi(groupStr)
		groupSizes = append(groupSizes, group)
	}
	return
}

func getNumArrangements(arrangement string, record string, groupSizes []int, currGroup int, inGroup bool) (sum int) {
	if len(record) == 0 {
		if currGroup == len(groupSizes) {
			return 1
		} else {
			return 0
		}
	}
	last := ""
	if len(arrangement) > 0 {
		last = string(arrangement[len(arrangement)-1])
	}
	if !(currGroup == len(groupSizes)) && (record[0] == '?' || record[0] == '#') && (last == "." || last == "" || (last == "#" && inGroup)) {
		if groupIsComplete(arrangement+"#", groupSizes[currGroup]) {
			sum += getNumArrangements(arrangement+"#", record[1:], groupSizes, currGroup+1, false)
		} else {
			sum += getNumArrangements(arrangement+"#", record[1:], groupSizes, currGroup, true)
		}
	}
	if !inGroup && (record[0] == '?' || record[0] == '.') {
		sum += getNumArrangements(arrangement+".", record[1:], groupSizes, currGroup, false)
	}
	return
}

func groupIsComplete(arrangement string, size int) bool {
	if len(arrangement) < size {
		return false
	}
	for _, symbol := range arrangement[len(arrangement)-(size):] {
		if symbol != '#' {
			return false
		}
	}
	return true
}
