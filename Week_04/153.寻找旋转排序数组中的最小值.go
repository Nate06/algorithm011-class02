/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	if nums[0] <= nums[len(nums) - 1] {
		return nums[0]
	}
	left := 0
	right := len(nums) - 1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid + 1] < nums[mid] {
			return nums[mid + 1]
		}
		if nums[mid - 1] > nums[mid] {
			return nums[mid]
		}

		if nums[0] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nums[mid]
}
// @lc code=end

