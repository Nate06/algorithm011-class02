/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	generate(len(nums), 0, nums, &res)
	return res
}

func generate(n, first int, nums []int, res *[][]int) {
	if first == n - 1 {
		*res = append(*res, append([]int{}, nums...))
		return
	}
	for i := first; i < len(nums); i++ {
		if i != first && nums[i] == nums[first] {
			continue
		}
		nums[first], nums[i] = nums[i], nums[first]
		generate(n, first + 1, nums, res)
	}
	for i := len(nums) - 1; i > first; i-- {
		nums[first], nums[i] = nums[i], nums[first]
	}
}
// @lc code=end

