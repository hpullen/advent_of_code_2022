package main

import (
	"day1/internal/buffer"
	"day1/internal/utils"
	"fmt"
)

func main() {
	data := utils.GetInput()[0]
	partOne(data)
	partTwo(data)
}

func partOne(data string) {
	fmt.Println("Part 1: ", findMarker(data, 4))
}

func partTwo(data string) {
	fmt.Println("Part 2: ", findMarker(data, 14))
}

func findMarker(data string, markerSize int) int {
	b := buffer.NewBuffer(markerSize)
	var n int
	for idx, c := range data {
		b.Push(c)
		if b.Len() == markerSize && b.AllUnique() {
			n = idx + 1
			break
		}
	}
	return n
}
