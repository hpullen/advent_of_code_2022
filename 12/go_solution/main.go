package main

import (
	"day1/internal/utils"
	"fmt"
)

func main() {
	data := utils.GetInput()
	points, start, summit := loadPoints(data)
	partOne(points, start, summit)
	partTwo(points, summit)
}

type point struct {
	i       int
	j       int
	value   rune
	score   int
	visited bool
}

func partOne(points [][]*point, start *point, summit *point) {
	steps := findSteps(points, start, summit)
	fmt.Println("Part 1: ", steps)
}

func partTwo(points [][]*point, summit *point) {
	starts := findPotentialStarts(points)
	currentMin := 10000
	for _, s := range starts {
		resetPoints(points, summit)
		score := findSteps(points, s, summit)
		if score < currentMin {
			currentMin = score
		}
	}
	fmt.Println("Part 2: ", currentMin)
}

func loadPoints(data []string) ([][]*point, *point, *point) {
	var points [][]*point
	var summit *point
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
				summit = &newPoint
			}
			pointLine = append(pointLine, &newPoint)
		}
		points = append(points, pointLine)
	}
	return points, start, summit
}

func findSteps(points [][]*point, start *point, summit *point) int {
	var priorityQueue []*point
	current := summit
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
		if len(priorityQueue) == 0 {
			// If we get here, there is no path between these points
			return 10000
		}
		current = priorityQueue[0]
		priorityQueue = priorityQueue[1:]
	}
	return current.score
}

func findPotentialStarts(points [][]*point) []*point {
	var starts []*point
	minVal := 'a'
	for _, p := range points[0] {
		if p.value == minVal {
			starts = append(starts, p)
		}
	}
	for _, p := range points[len(points)-1] {
		if p.value == minVal {
			starts = append(starts, p)
		}
	}
	for i := 1; i < len(points)-1; i++ {
		if points[i][0].value == minVal {
			starts = append(starts, points[i][0])
		}
		if points[i][len(points[0])-1].value == minVal {
			starts = append(starts, points[i][len(points[0])-1])
		}
	}
	return starts
}

func resetPoints(points [][]*point, summit *point) {
	for _, line := range points {
		for _, p := range line {
			p.score = 10000
			p.visited = false
		}
	}
	points[summit.i][summit.j].score = 0
}
