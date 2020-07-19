/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i - 1] {
			profit = prices[i] - prices[i - 1] + profit
		}
	}
	return profit
}
// @lc code=end

