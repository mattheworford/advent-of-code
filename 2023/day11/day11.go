package day11

import (
	"bufio"
	"math"
	"os"
	"slices"
)

func GetSumOfGalaxyDistances(documentName string) int {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	space := getSpace(scanner)
	galaxies, expandedCols, expandedRows := getGalaxies(space), getExpandedCols(space), getExpandedRows(space)
	distances := getDistances(galaxies, expandedCols, expandedRows)
	return sum(distances)
}

func getSpace(scanner *bufio.Scanner) (space []string) {
	for scanner.Scan() {
		line := scanner.Text()
		space = append(space, line)
	}
	return
}

func getExpandedCols(space []string) (cols []int) {
	expandable := map[int]bool{}
	for _, row := range space {
		for col, value := range row {
			_, exists := expandable[col]
			if !exists {
				expandable[col] = true
			}
			if value == '#' && expandable[col] {
				expandable[col] = false
			}
		}
	}
	var unsortedCols []int
	for col, isExpandable := range expandable {
		if isExpandable {
			unsortedCols = append(unsortedCols, col)
		}
	}
	slices.Sort(unsortedCols)
	for _, col := range unsortedCols {
		cols = append(cols, col)
	}
	return
}

func getExpandedRows(space []string) (rows []int) {
	for ind, row := range space {
		if !hasGalaxy(row) {
			rows = append(rows, ind)
		}
	}
	return
}

func hasGalaxy(row string) bool {
	for _, point := range row {
		if point == '#' {
			return true
		}
	}
	return false
}

func getGalaxies(space []string) (galaxies []Galaxy) {
	for x, row := range space {
		for y, point := range row {
			if point == '#' {
				galaxies = append(galaxies, Galaxy{
					x: x,
					y: y,
				})
			}
		}
	}
	return
}

func getDistances(galaxies []Galaxy, expandedCols, expandedRows []int) (distances []int) {
	for curr, first := range galaxies {
		for _, second := range galaxies[curr+1:] {
			maxX, minX := int(math.Max(float64(first.x), float64(second.x))), int(math.Min(float64(first.x), float64(second.x)))
			maxY, minY := int(math.Max(float64(first.y), float64(second.y))), int(math.Min(float64(first.y), float64(second.y)))
			xDistance := maxX - minX + len(filter(expandedRows, minX, maxX))
			yDistance := maxY - minY + len(filter(expandedCols, minY, maxY))
			distance := xDistance + yDistance
			distances = append(distances, distance)
		}
	}
	return
}

func filter(nums []int, min, max int) (filtered []int) {
	for _, num := range nums {
		if num >= min && num <= max {
			filtered = append(filtered, num)
		}
	}
	return
}

func sum(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}

type Galaxy struct {
	x int
	y int
}
