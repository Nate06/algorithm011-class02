/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int)  {
	k = k % len(nums)
	reverse := func(nums []int, from, to int) {
		for from < to {
			nums[from], nums[to] = nums[to], nums[from]
			from++
			to--
		}
	}
	reverse(nums, 0, len(nums) - 1)
	reverse(nums, 0, k - 1)
	reverse(nums, k, len(nums) - 1)
}
// @lc code=end

