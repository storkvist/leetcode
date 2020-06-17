package main

func solve(board [][]byte) {
	height := len(board)
	if height == 0 {
		return
	}
	width := len(board[0])
	if width == 0 {
		return
	}

	for i := 0; i < height; i++ {
		paint(board, width, height, i, 0)
		paint(board, width, height, i, width-1)
	}
	for j := 1; j < width-1; j++ {
		paint(board, width, height, 0, j)
		paint(board, width, height, height-1, j)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] == '!' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func paint(board [][]byte, width, height, i, j int) {
	if board[i][j] == 'O' {
		board[i][j] = '!'
		if i-1 >= 0 {
			paint(board, width, height, i-1, j)
		}
		if i+1 < height {
			paint(board, width, height, i+1, j)
		}
		if j-1 >= 0 {
			paint(board, width, height, i, j-1)
		}
		if j+1 < width {
			paint(board, width, height, i, j+1)
		}
	}
}
