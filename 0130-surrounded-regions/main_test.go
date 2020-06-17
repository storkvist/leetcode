package main

import "testing"

var tests = []struct {
	board [][]byte
	want  [][]byte
}{
	{
		[][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		},
		[][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'X', 'X'},
		},
	},
	{
		[][]byte{
			{},
			{},
		},
		[][]byte{
			{},
			{},
		},
	},
	{
		[][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'X', 'O'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'O', 'X'},
		},
		[][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'O'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'O', 'X'},
		},
	},
	{
		[][]byte{
			{'X', 'O', 'X', 'O', 'X', 'O'},
			{'O', 'X', 'O', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'O', 'X', 'O'},
			{'O', 'X', 'O', 'X', 'O', 'X'},
		},
		[][]byte{
			{'X', 'O', 'X', 'O', 'X', 'O'},
			{'O', 'X', 'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X', 'X', 'O'},
			{'O', 'X', 'O', 'X', 'O', 'X'},
		},
	},
	{
		[][]byte{
			{'O', 'X', 'X', 'O', 'X'},
			{'X', 'O', 'O', 'X', 'O'},
			{'X', 'O', 'X', 'O', 'X'},
			{'O', 'X', 'O', 'O', 'O'},
			{'X', 'X', 'O', 'X', 'O'},
		},
		[][]byte{
			{'O', 'X', 'X', 'O', 'X'},
			{'X', 'X', 'X', 'X', 'O'},
			{'X', 'X', 'X', 'O', 'X'},
			{'O', 'X', 'O', 'O', 'O'},
			{'X', 'X', 'O', 'X', 'O'},
		},
	},
}

func TestSolve(t *testing.T) {
	for _, test := range tests {
		board := make([][]byte, len(test.board))
		for i := range board {
			board[i] = make([]byte, len(test.board[i]))
			copy(board[i], test.board[i])
		}
		solve(board)
		if len(board) != len(test.want) {
			t.Errorf("solve(%v) -> %v; want %v\n", test.board, board, test.want)
			continue
		}
		for i := range board {
			ok := true
			if len(board[i]) != len(test.want[i]) {
				t.Errorf("solve(%v) -> %v; want %v\n", test.board, board, test.want)
				ok = false
			} else {
				for j := range board[i] {
					if board[i][j] != test.want[i][j] {
						t.Errorf("solve(%v) -> %v; want %v\n", test.board, board, test.want)
						ok = false
						break
					}
				}
			}
			if !ok {
				break
			}
		}
	}
}
