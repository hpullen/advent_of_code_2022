package main

import (
	"day1/internal/utils"
	"fmt"
	"sort"
	"strconv"
)

const (
	True = iota
	False
	Equal
)

func main() {
	data := utils.GetInput()
	pairs := getPairs(data)
	partOne(pairs)
	partTwo(pairs)
}

func partOne(pairs [][]string) {
	sum := 0
	for idx, p := range pairs {
		result := comparePair(p[0], p[1])
		if result == True {
			sum += idx + 1
		}
	}
	fmt.Println("Part 1:", sum)
}

func partTwo(pairs [][]string) {
	// Make slice containing all pairs to sort
	var allVals []string
	for _, p := range pairs {
		allVals = append(allVals, p[0])
		allVals = append(allVals, p[1])
	}
	div1 := "[[2]]"
	div2 := "[[6]]"
	allVals = append(allVals, div1)
	allVals = append(allVals, div2)

	// Sort
	sort.Slice(allVals, func(i, j int) bool {
		return comparePair(allVals[i], allVals[j]) == True
	})

	// Find indices of dividers
	var idx1, idx2 int
	for i, val := range allVals {
		if val == div1 {
			idx1 = i + 1
		} else if val == div2 {
			idx2 = i + 1
		}
	}
	fmt.Println("Part 2:", idx1*idx2)
}

func getPairs(data []string) [][]string {
	var pairs [][]string
	for i := 0; i < len(data); i += 3 {
		pairs = append(pairs, []string{data[i], data[i+1]})
	}
	return pairs
}

func comparePair(p1, p2 string) int {
	p1Split := split(p1)
	p2Split := split(p2)

	for i, p1Item := range p1Split {

		if i >= len(p2Split) {
			return False
		}
		p2Item := p2Split[i]

		p1IsList := p1Item[0] == '['
		p2IsList := p2Item[0] == '['

		if !p1IsList && !p2IsList {
			p1Val, _ := strconv.Atoi(p1Item)
			p2Val, _ := strconv.Atoi(p2Item)
			if p1Val < p2Val {
				return True
			} else if p1Val > p2Val {
				return False
			} else {
				continue
			}
		}

		if !p1IsList {
			p1Item = fmt.Sprintf("[%s]", p1Item)
		}
		if !p2IsList {
			p2Item = fmt.Sprintf("[%s]", p2Item)
		}
		result := comparePair(p1Item, p2Item)
		if result != Equal {
			return result
		}
	}

	if len(p2Split) > len(p1Split) {
		return True
	} else {
		return Equal
	}
}

func split(p string) []string {
	var items []string
	openBrackets := 0
	var currentItem string
	for _, c := range p[1 : len(p)-1] {
		if c == '[' {
			openBrackets += 1
		} else if c == ']' {
			openBrackets -= 1
		}
		if c == ',' && openBrackets == 0 {
			items = append(items, currentItem)
			currentItem = ""
		} else {
			currentItem = fmt.Sprintf("%s%s", currentItem, string(c))
		}
	}
	if len(currentItem) > 0 {
		items = append(items, currentItem)
	}
	return items
}
