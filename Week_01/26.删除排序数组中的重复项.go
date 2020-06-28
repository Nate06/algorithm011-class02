/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	i := 1
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[j - 1] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
// @lc code=end

