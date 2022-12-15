package gameRules

import "day1/internal/utils"

const (
	Rock = iota
	Paper
	Scissors
)

const (
	Win = iota
	Draw
	Loss
)

var winsAgainst = map[int]int{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

var losesAgainst = utils.ReverseMap(winsAgainst)

var scores = map[int]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

const winBonus = 6
const drawBonus = 3

func GetScoreFromMoves(theirMove, yourMove int) int {
	score := scores[yourMove]
	if winsAgainst[yourMove] == theirMove {
		score += winBonus
	} else if yourMove == theirMove {
		score += drawBonus
	}
	return score
}

func GetScoreFromOutcome(theirMove, outcome int) int {
	var yourMove int
	switch outcome {
	case Win:
		yourMove = losesAgainst[theirMove]
	case Draw:
		yourMove = theirMove
	case Loss:
		yourMove = winsAgainst[theirMove]
	}
	return GetScoreFromMoves(theirMove, yourMove)
}
