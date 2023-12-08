package day5

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func GetLowestLocationOfSeedRanges(documentName string) int {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	seeds := atoi(strings.Fields(strings.Split(scanner.Text(), ": ")[1]))
	seedRanges := getSeedRanges(seeds)
	scanner.Scan()
	maps := map[string]Map{}
	for scanner.Scan() {
		_, destination, m := parseMap(scanner)
		maps[destination] = m
	}
	for i := 0; true; i++ {
		seed := getSeed(i, maps)
		for _, seedRange := range seedRanges {
			if seedRange.contains(seed) {
				return i
			}
		}
	}
	return 0
}

func getSeedRanges(seeds []int) (seedRanges []SeedRange) {
	for i, seed := range seeds {
		if i%2 == 0 {
			seedRanges = append(seedRanges, SeedRange{
				start:  seed,
				length: seeds[i+1],
			})
		}
	}
	return
}

func getSeed(location int, maps map[string]Map) int {
	destination := "location"
	source := location
	for i := 0; i < 7; i++ {
		m, exists := maps[destination]
		if !exists {
			break
		}
		source = m.getSource(source)
		destination = m.source
	}
	return source
}

func GetLowestLocationOfSeeds(documentName string) int {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	seeds := atoi(strings.Fields(strings.Split(scanner.Text(), ": ")[1]))
	scanner.Scan()
	maps := map[string]Map{}
	for scanner.Scan() {
		source, _, m := parseMap(scanner)
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
	for i := 0; i < 7; i++ {
		m, exists := maps[source]
		if !exists {
			break
		}
		destination = m.getDestination(destination)
		source = m.destination
	}
	return destination
}

func parseMap(scanner *bufio.Scanner) (string, string, Map) {
	mapType := strings.Fields(scanner.Text())[0]
	source := strings.Split(mapType, "-")[0]
	destination := strings.Split(mapType, "-")[2]
	var ranges []MapRanges
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		} else {
			destRangeStart, _ := strconv.Atoi(strings.Fields(line)[0])
			sourceRangeStart, _ := strconv.Atoi(strings.Fields(line)[1])
			length, _ := strconv.Atoi(strings.Fields(line)[2])
			ranges = append(ranges, MapRanges{
				destRangeStart:   destRangeStart,
				sourceRangeStart: sourceRangeStart,
				length:           length,
			})
		}
	}
	return source, destination, Map{
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
	ranges      []MapRanges
}

func (m Map) getDestination(source int) int {
	for _, r := range m.ranges {
		if r.inSourceRange(source) {
			return r.destRangeStart + (source - r.sourceRangeStart)
		}
	}
	return source
}

func (m Map) getSource(destination int) int {
	for _, r := range m.ranges {
		if r.inDestinationRange(destination) {
			return r.sourceRangeStart + (destination - r.destRangeStart)
		}
	}
	return destination
}

type MapRanges struct {
	destRangeStart   int
	sourceRangeStart int
	length           int
}

func (r MapRanges) inSourceRange(source int) bool {
	fromStart := source - r.sourceRangeStart
	return 0 <= fromStart && fromStart <= r.length
}

func (r MapRanges) inDestinationRange(destination int) bool {
	fromStart := destination - r.destRangeStart
	return 0 <= fromStart && fromStart <= r.length
}

type SeedRange struct {
	start  int
	length int
}

func (r SeedRange) contains(seed int) bool {
	fromStart := seed - r.start
	return 0 <= fromStart && fromStart <= r.length
}
