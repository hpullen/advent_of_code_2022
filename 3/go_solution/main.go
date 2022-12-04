package main

import (
	"day1/internal/utils"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	data := utils.GetInput()
	partOne(data)
	partTwo(data)
}

func partOne(data []string) {
	total := 0
	for _, line := range data {
		compartmentSize := len(line) / 2
		c1 := line[0:compartmentSize]
		c2 := line[compartmentSize:]
		for _, c := range c1 {
			if strings.ContainsRune(c2, c) {
				total += getPriority(c)
				break
			}
		}
	}
	fmt.Println("Part 1: ", total)
}

func partTwo(data []string) {
	total := 0
	groupSize := 3
	for i := 0; i < len(data)/groupSize; i++ {
		firstElf := data[i*groupSize]
		for _, c := range firstElf {
			isAnswer := true
			for j := 1; j < groupSize; j++ {
				elf := data[i*groupSize+j]
				isAnswer = isAnswer && strings.ContainsRune(elf, c)
			}
			if isAnswer {
				total += getPriority(c)
				break
			}
		}
	}
	fmt.Println("Part 2: ", total)
}

func getPriority(c rune) int {
	if unicode.ToLower(c) == c {
		diff := c - 'a'
		return int(diff) + 1
	} else {
		diff := c - 'A'
		return int(diff) + 27
	}
}
