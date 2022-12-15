package main

import (
	"day1/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := utils.GetInput()
	dirs, steps := getSteps(data)
	// part_one(dirs, steps)
	part_two(dirs, steps)
}

func getSteps(data []string) ([]string, []int) {
	var dirs []string
	var steps []int
	for _, line := range data {
		line_split := strings.Split(line, " ")
		dirs = append(dirs, line_split[0])
		step, _ := strconv.Atoi(line_split[1])
		steps = append(steps, step)
	}
	return dirs, steps
}

type knot struct {
	x int
	y int
}

func part_one(dirs []string, steps []int) {
	visits := get_tail_visits(dirs, steps, 2)
	fmt.Println("Part 1:", visits)
}

func part_two(dirs []string, steps []int) {
	visits := get_tail_visits(dirs, steps, 10)
	fmt.Println("Part 2:", visits)
}

func get_tail_visits(dirs []string, steps []int, n int) int {
	var knots []knot
	for i := 0; i < n; i++ {
		knots = append(knots, knot{0, 0})
	}
	visited := make(map[string]bool)
	tail := len(knots) - 1
	log_visit(visited, knots[tail])
	for idx, dir := range dirs {
		for i := 0; i < steps[idx]; i++ {
			knots = take_step(dir, knots)
			log_visit(visited, knots[tail])
		}
	}
	return len(visited)
}

func take_step(dir string, knots []knot) []knot {
	switch dir {
	case "R":
		knots[0].x += 1
	case "L":
		knots[0].x -= 1
	case "U":
		knots[0].y += 1
	case "D":
		knots[0].y -= 1
	}

	for i := 1; i < len(knots); i++ {
		diff_x := knots[i-1].x - knots[i].x
		diff_y := knots[i-1].y - knots[i].y
		var dir_x int
		var dir_y int
		if diff_x != 0 {
			dir_x = diff_x / abs(diff_x)
		}
		if diff_y != 0 {
			dir_y = diff_y / abs(diff_y)
		}

		if abs(diff_x)+abs(diff_y) > 2 {
			knots[i].x += dir_x
			knots[i].y += dir_y
		} else {
			if abs(diff_x) == 2 {
				knots[i].x += dir_x
			} else if abs(diff_y) == 2 {
				knots[i].y += dir_y
			}
		}
	}
	return knots
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printPositions(knots []knot) {
	max_x := 0
	max_y := 0
	for _, k := range knots {
		if k.x > max_x {
			max_x = k.x
		}
		if k.y > max_y {
			max_y = k.y
		}
	}

	for y := max_y; y >= 0; y-- {
		var line string
		for x := 0; x <= max_x; x++ {
			toAdd := "."
			for idx := len(knots) - 1; idx >= 0; idx-- {
				if knots[idx].x == x && knots[idx].y == y {
					if idx == 0 {
						toAdd = "H"
					} else {
						toAdd = strconv.Itoa(idx)
					}
				}
			}
			line = fmt.Sprintf("%s%s", line, toAdd)
		}
		fmt.Println(line)
	}
}

func log_visit(visited map[string]bool, k knot) {
	key := fmt.Sprintf("%d,%d", k.x, k.y)
	visited[key] = true
}
