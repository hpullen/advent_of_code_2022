package main

import (
	"day1/internal/dirs"
	"day1/internal/utils"
	"fmt"
)

func main() {
	data := utils.GetInput()
	topDir, allDirs := dirs.FillDirs(data)
	partOne(allDirs)
	partTwo(topDir, allDirs)
}

func partOne(allDirs []*dirs.Dir) {
	total := 0
	const threshold = 100000
	for _, dir := range allDirs {
		if dir.Size <= threshold {
			total += dir.Size
		}
	}
	fmt.Println("Part 1: ", total)
}

func partTwo(topDir *dirs.Dir, allDirs []*dirs.Dir) {
	filled := topDir.Size
	const capacity = 70000000
	const spaceNeeded = 30000000
	const limit = capacity - spaceNeeded
	best := filled
	for _, dir := range allDirs {
		if (filled-dir.Size) <= limit && dir.Size < best {
			best = dir.Size
		}
	}
	fmt.Println("Part 2: ", best)
}
