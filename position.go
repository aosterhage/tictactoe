package main

import (
	"errors"
	"fmt"
	"math/bits"
)

type Side uint

const (
	// Crosses is first so that a default initialized Position has Crosses making the first move.
	Crosses Side = iota
	Noughts
)

type Result uint

const (
	Ongoing Result = iota
	Cat
	Win
)

// Bitboards use the lower 9-bits to represent the 3x3 board.
// The bits start at the lower left corner and move horizontally before moving vertically:
// 6 | 7 | 8
// ---------
// 3 | 4 | 5
// ---------
// 0 | 1 | 2
type Bitboard uint

type Position struct {
	noughts     Bitboard
	crosses     Bitboard
	active_side Side
}

var ErrTooManyPiecesProvided = errors.New("more than one piece provided in a single move")
var ErrPieceAlreadyAtLocation = errors.New("piece already exists at the provided location")

// Moves must only have a single bit set, corresponding to adding only a single piece at a time.
func (p *Position) MakeMove(b Bitboard) error {
	if bits.OnesCount(uint(b)) > 1 {
		return ErrTooManyPiecesProvided
	}

	if p.noughts&b > 0 || p.crosses&b > 0 {
		return ErrPieceAlreadyAtLocation
	}

	if p.active_side == Noughts {
		p.noughts |= b
		p.active_side = Crosses
	} else {
		p.crosses |= b
		p.active_side = Noughts
	}

	return nil
}

// The returned Side is only valid if the Result is "Win".
func (p Position) GetResult() (Result, Side) {
	wins := []Bitboard{
		// horizontal
		0x7,   // bottom row
		0x38,  // middle row
		0x1c0, // top row
		// vertical
		0x49,  // first column
		0x92,  // middle column
		0x124, // last column
		// diagonal
		0x111, // bottom left to top right
		0x54,  // top left to bottom right
	}

	for _, win := range wins {
		if win&p.noughts == win {
			return Win, Noughts
		} else if win&p.crosses == win {
			return Win, Crosses
		}
	}

	if p.noughts|p.crosses == 0x1ff {
		return Cat, Crosses
	}

	return Ongoing, Crosses
}

func (p Position) Print() {
	for y := 2; y >= 0; y-- {
		for x := 0; x <= 2; x++ {
			var b Bitboard = 1 << (3*y + x)
			if p.noughts&b > 0 {
				fmt.Printf(" O")
			} else if p.crosses&b > 0 {
				fmt.Printf(" X")
			} else {
				fmt.Printf("  ")
			}

			if x < 2 {
				fmt.Printf(" |")
			}
		}
		fmt.Printf("\n")

		if y > 0 {
			fmt.Println(" ---------")
		}
	}
	fmt.Println("")
}
