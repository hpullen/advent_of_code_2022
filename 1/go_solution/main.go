package main

import (
	"day1/internal/utils"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	calories := getSortedCalories()

	res1 := calories[0]
	fmt.Println("Part 1: ", res1)

	var top3 int
	for _, val := range calories[0:3] {
		top3 += val
	}
	fmt.Println("Part 2: ", top3)
}

func getSortedCalories() []int {
	rawCals := utils.GetInput()
	var elfCals []int
	var currentElf int
	for idx, cal := range rawCals {
		if cal != "" {
			calVal, err := strconv.Atoi(cal)
			utils.Check(err)
			currentElf += calVal
		}
		if cal == "" || idx == len(rawCals)-1 {
			elfCals = append(elfCals, currentElf)
			currentElf = 0
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(elfCals)))
	return elfCals
}
