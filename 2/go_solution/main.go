package main

import (
	"day1/internal/gameRules"
	"day1/internal/utils"
	"fmt"
	"strings"
)

var theirMoveMap = map[string]int{
	"A": gameRules.Rock,
	"B": gameRules.Paper,
	"C": gameRules.Scissors,
}

func main() {
	data := splitMoves(utils.GetInput())
	partOne(data)
	partTwo(data)
}

func splitMoves(data []string) [][]string {
	pairs := make([][]string, len(data))
	for idx, line := range data {
		pairs[idx] = strings.Split(line, " ")
	}
	return pairs
}

func partOne(data [][]string) {
	score := 0
	yourMoveMap := map[string]int{
		"X": gameRules.Rock,
		"Y": gameRules.Paper,
		"Z": gameRules.Scissors,
	}
	for _, moves := range data {
		theirMove := theirMoveMap[moves[0]]
		yourMove := yourMoveMap[moves[1]]
		score += gameRules.GetScoreFromMoves(theirMove, yourMove)
	}
	fmt.Println("Part 1: ", score)
}

func partTwo(data [][]string) {
	score := 0
	outcomesMap := map[string]int{
		"X": gameRules.Loss,
		"Y": gameRules.Draw,
		"Z": gameRules.Win,
	}
	for _, moves := range data {
		theirMove := theirMoveMap[moves[0]]
		outcome := outcomesMap[moves[1]]
		score += gameRules.GetScoreFromOutcome(theirMove, outcome)
	}
	fmt.Println("Part 2: ", score)
}
