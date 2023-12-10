package day10

import (
	"bufio"
	"os"
)

func GetStepsToFarthestPointInLoop(documentName string) (steps int) {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	var tiles []string
	for scanner.Scan() {
		tiles = append(tiles, scanner.Text())
	}
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
		return [][2]int{{i + 1, j}, {i, j - 1}}
	case 'F':
		return [][2]int{{i + 1, j}, {i, j + 1}}
	default:
		return [][2]int{}
	}
}
