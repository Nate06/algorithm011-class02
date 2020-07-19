/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	height := len(grid)
	if height == 0 {
		return 0
	}
	width := len(grid[0])
	islands := 0
	for x, line := range grid {
		for y, point := range line {
			if point == '1' {
				dfs(&grid, x, y, height, width)
				islands++
			}
		}
	}
	return islands
}

func dfs(grid *[][]byte, x, y, height, width int) {
	if x >= height || y >= width || x < 0 || y < 0 || (*grid)[x][y] == '0' {
		return
	}
	(*grid)[x][y] = '0'
	dfs(grid, x + 1, y, height, width)
	dfs(grid, x - 1, y, height, width)
	dfs(grid, x, y + 1, height, width)
	dfs(grid, x, y - 1, height, width)
}
// @lc code=end

