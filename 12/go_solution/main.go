package main

import (
	"day1/internal/utils"
	"fmt"
)

func main() {
	data := utils.GetInput()
	partOne(data)
}

type point struct {
	i       int
	j       int
	value   rune
	score   int
	visited bool
}

func partOne(data []string) {
	// Load the map
	var points [][]*point
	var current *point
	var start *point
	for ii, line := range data {
		var pointLine []*point
		for jj, c := range line {
			newPoint := point{ii, jj, c, 100000, false}
			if c == 'S' {
				newPoint.value = 'a'
				start = &newPoint
			} else if c == 'E' {
				newPoint.value = 'z'
				newPoint.score = 0
				current = &newPoint
			}
			pointLine = append(pointLine, &newPoint)
		}
		points = append(points, pointLine)
	}

	var priorityQueue []*point

	// Loop until we reach the start
	for current != start {
		current.visited = true

		// Update neighbours
		var neighbours []*point
		i := current.i
		j := current.j
		if i > 0 {
			neighbours = append(neighbours, points[i-1][j])
		}
		if i < len(points)-1 {
			neighbours = append(neighbours, points[i+1][j])
		}
		if j > 0 {
			neighbours = append(neighbours, points[i][j-1])
		}
		if j < len(points[0])-1 {
			neighbours = append(neighbours, points[i][j+1])
		}

		for _, n := range neighbours {
			if current.value-n.value <= 1 {
				if current.score+1 < n.score {
					n.score = current.score + 1
					priorityQueue = append(priorityQueue, n)
				}
			}
		}
		current = priorityQueue[0]
		priorityQueue = priorityQueue[1:]
	}
	fmt.Println("Part 1: ", current.score)
}
