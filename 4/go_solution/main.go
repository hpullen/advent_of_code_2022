package main

import (
	"day1/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := getRanges(utils.GetInput())
	partOne(data)
	partTwo(data)
}

func getRanges(data []string) [][][]int {
	var ranges [][][]int
	for _, line := range data {
		var current [][]int
		pair := strings.Split(line, ",")
		for _, elf := range pair {
			rangeStrs := strings.Split(elf, "-")
			var currentRange []int
			for _, valStr := range rangeStrs {
				val, _ := strconv.Atoi(valStr)
				currentRange = append(currentRange, val)
			}
			current = append(current, currentRange)
		}
		ranges = append(ranges, current)
	}
	return ranges
}

func partOne(data [][][]int) {
	count := 0
	for _, ranges := range data {
		fullyContained := false
		for i := 0; i <= 1; i++ {
			j := 1 - i
			if ranges[i][0] <= ranges[j][0] && ranges[i][1] >= ranges[j][1] {
				fullyContained = true
				break
			}
		}
		if fullyContained {
			count += 1
		}
	}
	fmt.Println("Part 1: ", count)
}

func partTwo(data [][][]int) {
	count := 0
	for _, ranges := range data {
		overlap := false
		for i := 0; i <= 1; i++ {
			for j := 0; j <= 1; j++ {
				if ranges[j][i] >= ranges[1-j][0] && ranges[j][i] <= ranges[1-j][1] {
					overlap = true
					break
				}
			}
		}
		if overlap {
			count += 1
		}
	}
	fmt.Println("Part 2: ", count)
}
