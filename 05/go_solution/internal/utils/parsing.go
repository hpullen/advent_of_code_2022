package utils

import (
	"day1/internal/stack"
	"strconv"
	"strings"
)

func GetStacks(data []string) []stack.Stack {
	var endOfStacksIdx int
	nStacks := 0
	for idx, line := range data {
		if strings.Contains(line, " 1 ") {
			endOfStacksIdx = idx
			lineContents := strings.Split(line, " ")
			for _, c := range lineContents {
				if c != "" {
					nStacks += 1
				}
			}
			break
		}
	}

	var stacks []stack.Stack
	for i := 0; i < nStacks; i++ {
		stacks = append(stacks, []string{})
	}
	for i := 0; i < endOfStacksIdx; i++ {
		row := data[endOfStacksIdx-1-i]
		j := 0
		n := 0
		for j < len(row) {
			item := row[j : j+3]
			if item != "   " {
				stacks[n].Push(string(item[1]))
			}
			j += 4
			n += 1
		}
	}
	return stacks
}

type Move struct {
	N     int
	Start int
	End   int
}

func GetMoves(data []string) []Move {
	var moves []Move
	for _, line := range data {
		if !strings.Contains(line, "move") {
			continue
		}
		words := strings.Split(line, " ")
		N, _ := strconv.Atoi(words[1])
		Start, _ := strconv.Atoi(words[3])
		End, _ := strconv.Atoi(words[5])
		moves = append(moves, Move{N, Start, End})
	}
	return moves
}
