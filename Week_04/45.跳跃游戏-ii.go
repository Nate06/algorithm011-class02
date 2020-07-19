/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	maxPosition := 0
	step := 0
	end := 0
	for i := 0; i < len(nums) - 1; i++ {
		if i + nums[i] > maxPosition {
			maxPosition = i + nums[i]
		}
		if i == end {
			end = maxPosition
			step++
		}
	}
	return step
}
// @lc code=end

