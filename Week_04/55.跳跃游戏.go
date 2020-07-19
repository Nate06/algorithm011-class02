/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	rightMost := 0
	for i := 0; i < len(nums); i++ {
		if i <= rightMost {
			if i + nums[i] > rightMost {
				rightMost = i + nums[i]
			}
			if rightMost >= len(nums) - 1 {
				return true
			}
		}
	}
	return false
}
// @lc code=end

