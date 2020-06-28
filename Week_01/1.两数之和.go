/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	numsMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		searchVal := target - nums[i]
		if index, ok := numsMap[searchVal]; ok {
			return []int{index, i}
		}
		numsMap[nums[i]] = i
	}
	return []int{}
}
// @lc code=end

