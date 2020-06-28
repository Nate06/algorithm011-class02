/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	i := 0
	j := len(height) - 1
	leftMax, rightMax, sum := 0, 0, 0
	for i < j {
		if height[i] < height[j] {
			if height[i] < leftMax {
				sum += leftMax - height[i]
			} else {
				leftMax = height[i]
			}
			i++
		} else {
			if height[j] < rightMax {
				sum += rightMax - height[j]
			} else {
				rightMax = height[j]
			}
			j--
		}
	}
	return sum
}
// @lc code=end

