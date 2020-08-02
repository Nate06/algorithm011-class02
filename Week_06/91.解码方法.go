/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, cur, tmp := 1, 1, 0 
	for i := 1; i < len(s); i++ {
		if s[i] == '0' && s[i - 1] != '1' && s[i - 1] != '2' {
			return 0
		} else if s[i] == '0' {
			cur = pre
		} else if s[i - 1] == '1' || (s[i - 1] == '2' && s[i] <= '6') {
			tmp = cur 
			cur = cur + pre
			pre = tmp
		} else {
			pre = cur
		}
	}
	return cur
}
// @lc code=end

