/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	numsMap := map[int]int{}
	for i, num := range nums {
		anotherNum := target - num
		if j, ok := numsMap[anotherNum]; ok {
			return []int{i, j}
		}
		numsMap[num] = i
	}
	return []int{}
}
// @lc code=end

