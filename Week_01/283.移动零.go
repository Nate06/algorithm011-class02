/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int)  {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			temp := nums[j]
			nums[j] = nums[i]
			nums[i] = temp
			i++
		}
	}
}
// @lc code=end

