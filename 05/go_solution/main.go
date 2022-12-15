package main

import (
	"day1/internal/stack"
	"day1/internal/utils"
	"fmt"
)

func main() {
	data := utils.GetInput()
	partOne(data)
	partTwo(data)
}

func partOne(data []string) {
	fmt.Println("Part 1: ", solve(data, applyMoves1))
}

func partTwo(data []string) {
	fmt.Println("Part 2: ", solve(data, applyMoves2))
}

func solve(data []string, moverFunc func([]stack.Stack, []utils.Move)) string {
	stacks := utils.GetStacks(data)
	moves := utils.GetMoves(data)
	moverFunc(stacks, moves)
	var answer string
	for _, stack := range stacks {
		answer = fmt.Sprintf("%s%s", answer, stack.Pop())
	}
	return answer
}

func applyMoves1(stacks []stack.Stack, moves []utils.Move) {
	for _, move := range moves {
		for i := 0; i < move.N; i++ {
			stacks[move.End-1].Push(stacks[move.Start-1].Pop())
		}
	}
}

func applyMoves2(stacks []stack.Stack, moves []utils.Move) {
	for _, move := range moves {
		stacks[move.End-1].PushMany(stacks[move.Start-1].PopMany(move.N))
	}
}