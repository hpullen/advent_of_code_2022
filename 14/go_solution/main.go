package main

import (
	"day1/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	coords := getCoords(utils.GetInput())
	partOne(coords)
	partTwo(coords)
}

func partOne(coords [][]coord) {
	caveMap, start := makeMap(coords, false)
	count := 0
	for {
		_, fellIntoVoid := simulateSingleGrain(caveMap, start)
		if fellIntoVoid {
			break
		}
		count += 1
	}
	fmt.Println("Part 1:", count)
}

func partTwo(coords [][]coord) {
	caveMap, start := makeMap(coords, true)
	count := 0
	for {
		count += 1
		finalCoord, _ := simulateSingleGrain(caveMap, start)
		if finalCoord.i == 0 && finalCoord.j == start {
			break
		}
	}
	fmt.Println("Part 2:", count)
}

func simulateSingleGrain(caveMap [][]bool, start int) (coord, bool) {
	j := start
	for i := 0; i < len(caveMap); i++ {
		// Keep falling if the space below is clear
		if !caveMap[i+1][j] {
			continue
		}

		// Try going left
		if j == 0 {
			return coord{}, true
		}
		if !caveMap[i+1][j-1] {
			j -= 1
			continue
		}

		// Try going right
		if j == len(caveMap[0])-1 {
			return coord{}, true
		}
		if !caveMap[i+1][j+1] {
			j += 1
			continue
		}

		// If we get here, the sand comes to rest
		caveMap[i][j] = true
		return coord{i, j}, false
	}
	return coord{}, true
}

type coord struct {
	i, j int
}

func getCoords(data []string) [][]coord {
	var coords [][]coord
	for _, line := range data {
		var row []coord
		for _, point := range strings.Split(line, " -> ") {
			coordInput := strings.Split(point, ",")
			i, _ := strconv.Atoi(coordInput[1])
			j, _ := strconv.Atoi(coordInput[0])
			row = append(row, coord{i, j})
		}
		coords = append(coords, row)
	}
	return coords
}

func makeMap(coords [][]coord, wide bool) ([][]bool, int) {
	minCoord, maxCoord := getRange(coords)
	source_j := 500
	if wide {
		maxCoord.i += 2
		minCoord.j = source_j - maxCoord.i - 2
		maxCoord.j = source_j + maxCoord.i + 2
	}
	start := source_j - minCoord.j

	// Make empty map
	var caveMap [][]bool
	for i := minCoord.i; i <= maxCoord.i; i++ {
		row := make([]bool, maxCoord.j+1-minCoord.j)
		caveMap = append(caveMap, row)
	}

	// Fill based on coordinates
	for _, rocks := range coords {
		for idx := 0; idx < len(rocks)-1; idx++ {
			start := rocks[idx]
			end := rocks[idx+1]
			var start_i, start_j, end_i, end_j int
			if start.i < end.i {
				start_i = start.i
				end_i = end.i
			} else {
				start_i = end.i
				end_i = start.i
			}
			if start.j < end.j {
				start_j = start.j
				end_j = end.j
			} else {
				start_j = end.j
				end_j = start.j
			}
			for i := start_i; i <= end_i; i++ {
				for j := start_j; j <= end_j; j++ {
					caveMap[i][j-minCoord.j] = true
				}
			}
		}
	}

	// Fill floor
	if wide {
		for j := 0; j < len(caveMap[0]); j++ {
			caveMap[len(caveMap)-1][j] = true
		}
	}

	return caveMap, start
}

func getRange(coords [][]coord) (coord, coord) {
	minCoord := coord{0, 10000}
	maxCoord := coord{0, 0}
	for _, row := range coords {
		for _, point := range row {
			if point.i > maxCoord.i {
				maxCoord.i = point.i
			}
			if point.j < minCoord.j {
				minCoord.j = point.j
			}
			if point.j > maxCoord.j {
				maxCoord.j = point.j
			}
		}
	}
	return minCoord, maxCoord
}

func printMap(caveMap [][]bool, start int) {
	for i, row := range caveMap {
		toPrint := ""
		for j, point := range row {
			if i == 0 && j == start {
				toPrint = fmt.Sprintf("%s+", toPrint)
			} else if point {
				toPrint = fmt.Sprintf("%s#", toPrint)
			} else {
				toPrint = fmt.Sprintf("%s.", toPrint)
			}
		}
		fmt.Println(toPrint)
	}
}
