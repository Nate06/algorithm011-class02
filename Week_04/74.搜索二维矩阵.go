/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	height := len(matrix)
	if height == 0 {
		return false
	}
	width := len(matrix[0])

	left := 0
	right := height * width - 1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		x := mid / width
		y := mid % width
		current := matrix[x][y]
		if current == target {
			return true
		}
		if target < current {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
// @lc code=end

