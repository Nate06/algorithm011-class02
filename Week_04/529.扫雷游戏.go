/*
 * @lc app=leetcode.cn id=529 lang=golang
 *
 * [529] 扫雷游戏
 */

// @lc code=start
func updateBoard(board [][]byte, click []int) [][]byte {
	height := len(board)
	if height == 0 {
		return board
	}
	width := len(board[0])

	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	positions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}
	update(&board, click[0], click[1], height, width, positions)

	return board
}

func update(board *[][]byte, x, y, height, width int, positions [][]int) {
	if x >= height || x < 0 || y >= width || y < 0 || (*board)[x][y] != 'E' {
		return
	}

	mine := 0
	for _, position := range positions {
		if x + position[0] < height && x + position[0] >= 0 && y + position[1] < width && y + position[1] >= 0 {
			if (*board)[x + position[0]][y + position[1]] == 'M' {
				mine++
			}
		}
	}
	if mine > 0 {
		(*board)[x][y] = byte('0' + mine)
	} else {
		(*board)[x][y] = 'B'
		for _, position := range positions {
			update(board, x + position[0], y + position[1], height, width, positions)
		}
	}
}
// @lc code=end

