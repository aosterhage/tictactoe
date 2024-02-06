package main

import (
	"testing"
)

func TestMakeMove(t *testing.T) {
	type test struct {
		name            string
		position        Position
		move            Bitboard
		expect_error    error
		expect_position Position
	}

	tests := []test{
		{
			// start with an empty board and fill the first square (bottom left)
			name:            "ValidFirstMove",
			position:        Position{noughts: 0x0, crosses: 0x0, active_side: Crosses},
			move:            0x1,
			expect_error:    nil,
			expect_position: Position{noughts: 0x0, crosses: 0x1, active_side: Noughts},
		},
		{
			// try to place a piece on a square that already contains a piece
			name:            "ErrorPieceAlreadyAtLocation",
			position:        Position{noughts: 0x0, crosses: 0x1, active_side: Noughts},
			move:            0x1,
			expect_error:    ErrPieceAlreadyAtLocation,
			expect_position: Position{noughts: 0x0, crosses: 0x1, active_side: Noughts},
		},
		{
			// try to place multiple pieces
			name:            "ErrorTooManyPiecesProvided",
			position:        Position{noughts: 0x0, crosses: 0x0, active_side: Crosses},
			move:            0x3,
			expect_error:    ErrTooManyPiecesProvided,
			expect_position: Position{noughts: 0x0, crosses: 0x0, active_side: Crosses},
		},
		{
			// valid move from a random position
			name:            "ValidRandomPosition",
			position:        Position{noughts: 0x10, crosses: 0x101, active_side: Noughts},
			move:            0x4,
			expect_error:    nil,
			expect_position: Position{noughts: 0x14, crosses: 0x101, active_side: Crosses},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.position.MakeMove(tt.move)

			if err != tt.expect_error {
				t.Errorf("\ngot error = %v\nexpected = %v", err, tt.expect_error)
			} else if tt.position != tt.expect_position {
				t.Errorf("\ngot position = %v\nexpected = %v", tt.position, tt.expect_position)
			}
		})
	}
}

func TestGetResult(t *testing.T) {
	type test struct {
		name          string
		position      Position
		expect_result Result
		expect_winner Side
	}

	tests := []test{
		{
			// start position
			name:          "StartPosition",
			position:      Position{noughts: 0x0, crosses: 0x0, active_side: Crosses},
			expect_result: Ongoing,
		},
		{
			// cat game
			name:          "CatGame",
			position:      Position{noughts: 0xf1, crosses: 0x10e, active_side: Crosses},
			expect_result: Cat,
		},
		{
			// noughts win
			name:          "NoughtsWin",
			position:      Position{noughts: 0x7, crosses: 0x58, active_side: Crosses},
			expect_result: Win,
			expect_winner: Noughts,
		},
		{
			// crosses win
			name:          "CrossesWin",
			position:      Position{noughts: 0x58, crosses: 0x7, active_side: Noughts},
			expect_result: Win,
			expect_winner: Crosses,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, winner := tt.position.GetResult()

			if result != tt.expect_result {
				t.Errorf("\ngot result = %v\nexpected = %v", result, tt.expect_result)
			} else if result == Win && winner != tt.expect_winner {
				t.Errorf("\ngot winner = %v\nexpected = %v", winner, tt.expect_winner)
			}
		})
	}
}
