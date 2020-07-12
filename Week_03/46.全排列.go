/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := [][]int{}
	generate(len(nums), 0, nums, &res)
	return res
}

func generate(n, first int, nums []int, res *[][]int) {
	if first == n {
		*res = append(*res, append([]int{}, nums...))
		return
	}
	for i := first; i < len(nums); i++ {
		nums[i], nums[first] = nums[first], nums[i]
		generate(n, first + 1, nums, res)
		nums[i], nums[first] = nums[first], nums[i]
	}
}
// @lc code=end

