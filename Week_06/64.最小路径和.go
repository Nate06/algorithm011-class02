/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	res := [][]int{}
	for i := range grid {
		res = append(res, []int{})
		for _ = range grid[0] {
			res[i] = append(res[i], 0)
		}
	}
	for x, row := range grid {
		for y, num := range row {
			if x == 0 && y == 0 {
				res[x][y] = num
			} else if x == 0 {
				res[x][y] = num + res[x][y - 1]
			} else if y == 0 {
				res[x][y] = num + res[x - 1][y]
			} else {
				if res[x - 1][y] < res[x][y - 1] {
					res[x][y] = res[x - 1][y] + num
				} else {
					res[x][y] = num + res[x][y - 1]
				}
			}
		}
	}
	return res[len(res) - 1][len(res[0]) - 1]
}
// @lc code=end

