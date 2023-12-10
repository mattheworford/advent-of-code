package day10

import (
	"bufio"
	"os"
)

func GetEnclosedTiles(documentName string) (enclosed int) {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	tiles := getTiles(scanner)
	start, pipeMap := getPipeMap(tiles)
	loopLength := getLoopLength(start, pipeMap)
	for i := 0; i < 2; i++ {
		corners := getCorners(start, pipeMap[start][i], pipeMap)
		enclosed = getInteriorArea(loopLength, corners)
		if enclosed >= 0 {
			break
		}
	}
	return
}

func GetStepsToFarthestPointInLoop(documentName string) (steps int) {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	tiles := getTiles(scanner)
	start, pipeMap := getPipeMap(tiles)
	lastLeft, lastRight, left, right := start, start, pipeMap[start][0], pipeMap[start][1]
	steps += 1
	for left != right {
		nextLeft, nextRight := getNext(lastLeft, left, pipeMap), getNext(lastRight, right, pipeMap)
		lastLeft, lastRight, left, right = left, right, nextLeft, nextRight
		steps += 1
	}
	return
}

func getCorners(start, first [2]int, pipeMap map[[2]int][2][2]int) (corners [][2]int) {
	last, curr := start, first
	for curr != start {
		next := getNext(last, curr, pipeMap)
		last, curr = curr, next
		if isCorner(pipeMap[curr]) {
			corners = append(corners, curr)
		}
	}
	return
}

func getLoopLength(start [2]int, pipeMap map[[2]int][2][2]int) int {
	length, prev, curr := 1, start, pipeMap[start][1]
	for curr != start {
		length += 1
		next := getNext(prev, curr, pipeMap)
		prev, curr = curr, next
	}
	return length
}

func getInteriorArea(perimeter int, corners [][2]int) (area int) {
	return getTotalArea(corners) + 1 - (perimeter / 2)
}

func getTotalArea(corners [][2]int) (area int) {
	for i, corner := range corners {
		if i < len(corners)-1 {
			area += (corner[0] * corners[i+1][1]) - (corners[i+1][0] * corner[1])
		} else {
			area += (corner[0] * corners[0][1]) - (corners[0][0] * corner[1])
		}
	}
	return area / 2
}

func isCorner(neighbors [2][2]int) bool {
	left, right := neighbors[0], neighbors[1]
	if left[0] == right[0] || left[1] == right[1] {
		return false
	}
	return true
}

func getTiles(scanner *bufio.Scanner) (tiles []string) {
	for scanner.Scan() {
		line := scanner.Text()
		tiles = append(tiles, line)
	}
	return
}

func getPipeMap(tiles []string) ([2]int, map[[2]int][2][2]int) {
	start, coordMap := [2]int{}, map[[2]int][2][2]int{}
	for i := range tiles {
		for j, tile := range tiles[i] {
			if tile == 'S' {
				start = [2]int{i, j}
				coordMap[[2]int{i, j}] = [2][2]int(parseStart(start, tiles))
			} else {
				neighbors := getNeighbors(i, j, tile)
				if len(neighbors) > 0 {
					coordMap[[2]int{i, j}] = [2][2]int(neighbors)
				}
			}
		}
	}
	return start, coordMap
}

func getNext(last, curr [2]int, coordMap map[[2]int][2][2]int) [2]int {
	neighbors, _ := coordMap[curr]
	for _, neighbor := range neighbors {
		if neighbor != last {
			return neighbor
		}
	}
	return [2]int{}
}

func parseStart(coords [2]int, tiles []string) (res [][2]int) {
	if coords[0] > 0 {
		north := [2]int{coords[0] - 1, coords[1]}
		if areNeighbors(coords, north, tiles) {
			res = append(res, north)
		}
	}
	if coords[1] > 0 {
		west := [2]int{coords[0], coords[1] - 1}
		if areNeighbors(coords, west, tiles) {
			res = append(res, west)
		}
	}
	if coords[0] < len(tiles)-1 {
		south := [2]int{coords[0] + 1, coords[1]}
		if areNeighbors(coords, south, tiles) {
			res = append(res, south)
		}
	}
	if coords[1] < len(tiles[coords[0]])-1 {
		east := [2]int{coords[0], coords[1] + 1}
		if areNeighbors(coords, east, tiles) {
			res = append(res, east)
		}
	}
	return
}

func areNeighbors(first, second [2]int, tiles []string) bool {
	x, y := second[0], second[1]
	neighbors := getNeighbors(x, y, int32(tiles[x][y]))
	for i := range neighbors {
		if neighbors[i] == first {
			return true
		}
	}
	return false
}

func getNeighbors(i, j int, tile int32) [][2]int {
	switch tile {
	case '|':
		return [][2]int{{i - 1, j}, {i + 1, j}}
	case '-':
		return [][2]int{{i, j - 1}, {i, j + 1}}
	case 'L':
		return [][2]int{{i - 1, j}, {i, j + 1}}
	case 'J':
		return [][2]int{{i - 1, j}, {i, j - 1}}
	case '7':
		return [][2]int{{i, j - 1}, {i + 1, j}}
	case 'F':
		return [][2]int{{i, j + 1}, {i + 1, j}}
	default:
		return [][2]int{}
	}
}
