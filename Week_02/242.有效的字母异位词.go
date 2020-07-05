/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counter := make([]int, 26)
	for i := range s {
		counter[s[i] - 'a']++
		counter[t[i] - 'a']--
	}

	for _, c := range counter {
		if c != 0 {
			return false
		}
	}
	return true
}
// @lc code=end

