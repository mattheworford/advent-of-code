package day5

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func GetLowestLocation(documentName string) int {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	seeds := atoi(strings.Fields(strings.Split(scanner.Text(), ": ")[1]))
	fmt.Println(seeds)
	scanner.Scan()
	maps := map[string]Map{}
	for scanner.Scan() {
		source, m := parseMap(scanner)
		maps[source] = m
	}
	lowestLocation := math.MaxInt
	for _, seed := range seeds {
		lowestLocation = min(lowestLocation, getLocation(seed, maps))
	}
	return lowestLocation
}

func getLocation(seed int, maps map[string]Map) int {
	source := "seed"
	destination := seed
	for {
		m, exists := maps[source]
		if !exists {
			break
		}
		destination = m.getDestination(destination)
		source = m.destination
	}
	return destination
}

func (m Map) getDestination(source int) int {
	for _, r := range m.ranges {
		if r.inSourceRange(source) {
			return r.destRangeStart + (source - r.sourceRangeStart)
		}
	}
	return source
}

func (r Ranges) inSourceRange(source int) bool {
	fromStart := source - r.sourceRangeStart
	return 0 <= fromStart && fromStart <= r.length
}

func parseMap(scanner *bufio.Scanner) (string, Map) {
	mapType := strings.Fields(scanner.Text())[0]
	source := strings.Split(mapType, "-")[0]
	destination := strings.Split(mapType, "-")[2]
	var ranges []Ranges
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		} else {
			destRangeStart, _ := strconv.Atoi(strings.Fields(line)[0])
			sourceRangeStart, _ := strconv.Atoi(strings.Fields(line)[1])
			length, _ := strconv.Atoi(strings.Fields(line)[2])
			ranges = append(ranges, Ranges{
				destRangeStart:   destRangeStart,
				sourceRangeStart: sourceRangeStart,
				length:           length,
			})
		}
	}
	return source, Map{
		source:      source,
		destination: destination,
		ranges:      ranges,
	}
}

func atoi(strings []string) (ints []int) {
	for _, str := range strings {
		num, _ := strconv.Atoi(str)
		ints = append(ints, num)
	}
	return
}

type Map struct {
	source      string
	destination string
	ranges      []Ranges
}

type Ranges struct {
	destRangeStart   int
	sourceRangeStart int
	length           int
}
