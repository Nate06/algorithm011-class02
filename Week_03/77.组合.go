/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := [][]int{}
	if n < k {
		return res
	}
	generate(n, k, 1, []int{}, &res)
	return res
}

func generate(n, k, first int, nums []int, res *[][]int) {
	if len(nums) == k {
		*res = append(*res, append([]int{}, nums...))
		return
	}
	for i := first; i <= n; i++ {
		nums = append(nums, i)
		generate(n, k, i + 1, nums, res)
		nums = nums[:len(nums) - 1]
	}
}
// @lc code=end

