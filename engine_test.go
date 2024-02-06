package main

import (
	"reflect"
	"testing"
)

func TestGenerateLegalMoves(t *testing.T) {
	type test struct {
		name     string
		position Position
		expect   []Bitboard
	}

	tests := []test{
		{
			name:     "EmptyBoard",
			position: Position{noughts: 0x0, crosses: 0x0, active_side: Crosses},
			expect: []Bitboard{
				0x1, 0x2, 0x4, 0x8, 0x10, 0x20, 0x40, 0x80, 0x100,
			},
		},
		{
			name:     "RandomBoard",
			position: Position{noughts: 0x10, crosses: 0x101, active_side: Noughts},
			expect: []Bitboard{
				0x2, 0x4, 0x8, 0x20, 0x40, 0x80,
			},
		},
		{
			name:     "CatGame",
			position: Position{noughts: 0xf1, crosses: 0x10e, active_side: Crosses},
			expect:   []Bitboard{},
		},
		{
			name:     "NoughtsWin",
			position: Position{noughts: 0x7, crosses: 0x58, active_side: Crosses},
			expect:   []Bitboard{},
		},
		{
			name:     "CrossesWin",
			position: Position{noughts: 0x58, crosses: 0x7, active_side: Noughts},
			expect:   []Bitboard{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moves := GenerateLegalMoves(tt.position)

			if !reflect.DeepEqual(moves, tt.expect) {
				t.Errorf("\ngot moves = %v\nexpected = %v", moves, tt.expect)
			}
		})
	}
}

func TestEvalMove(t *testing.T) {
	type test struct {
		name     string
		position Position
		move     Bitboard
		expect   int
	}

	tests := []test{
		{
			name:     "Cat",
			position: Position{noughts: 0x11a, crosses: 0xc5, active_side: Crosses},
			move:     0x20,
			expect:   0,
		},
		{
			name:     "InstantWin",
			position: Position{noughts: 0x181, crosses: 0x4e, active_side: Crosses},
			move:     0x10,
			expect:   10,
		},
		{
			name:     "InstantLoss",
			position: Position{noughts: 0x181, crosses: 0x4e, active_side: Crosses},
			move:     0x20,
			expect:   -10,
		},
		{
			name:     "Win",
			position: Position{noughts: 0x0, crosses: 0x44, active_side: Crosses},
			move:     0x1,
			expect:   1,
		},
		{
			name:     "Loss",
			position: Position{noughts: 0x0, crosses: 0x44, active_side: Noughts},
			move:     0x10,
			expect:   -1,
		},
		{
			name:     "StartPosition",
			position: Position{noughts: 0x0, crosses: 0x0, active_side: Crosses},
			move:     0x1,
			expect:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eval := EvalMove(0, tt.position.active_side, tt.position, tt.move)
			if eval != tt.expect {
				t.Errorf("\ngot eval = %v\nexpected = %v", eval, tt.expect)
			}
		})
	}
}
