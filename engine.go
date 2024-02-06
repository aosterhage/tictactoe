package main

import (
	"math/rand"
	"slices"
)

func GenerateLegalMoves(p Position) []Bitboard {
	var moves = []Bitboard{}

	result, _ := p.GetResult()
	if result == Ongoing {
		all_pieces := p.noughts | p.crosses
		for i := 0; i <= 8; i++ {
			var b Bitboard = 1 << i
			if b&all_pieces == 0 {
				moves = append(moves, b)
			}
		}
	}

	return moves
}

func GetRandomMove(p Position) Bitboard {
	moves := GenerateLegalMoves(p)
	return moves[rand.Intn(len(moves))]
}

func EvalMove(depth int, s Side, p Position, b Bitboard) int {
	p.MakeMove(b)
	result, winner := p.GetResult()

	// base cases
	if result == Cat {
		return 0
	} else if result == Win {
		value := 1
		if depth < 2 {
			value *= 10
		}

		if winner == s {
			return value
		} else {
			return -value
		}
	}

	// otherwise we need to go down one depth
	var evals []int

	if result == Ongoing {
		moves := GenerateLegalMoves(p)
		for _, move := range moves {
			evals = append(evals, EvalMove(depth+1, s, p, move))
		}
	}

	if p.active_side == s {
		return slices.Max(evals)
	} else {
		return slices.Min(evals)
	}
}

func GetBestMove(p Position) Bitboard {
	moves := GenerateLegalMoves(p)
	var evals []int
	for _, move := range moves {
		evals = append(evals, EvalMove(0, p.active_side, p, move))
	}

	var best_move_index int = 0
	for i, eval := range evals {
		if eval > evals[best_move_index] {
			best_move_index = i
		}
	}

	return moves[best_move_index]
}
