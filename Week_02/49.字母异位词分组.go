/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	resMap := map[string][]string{}
	for _, str := range strs {
		counter := make([]int, 26)
		for _, s := range str {
			counter[s - 'a']++
		}
		counterStr := ""
		for _, c := range counter {
			counterStr += string(c) + "#"
		}
		if group, ok := resMap[counterStr]; ok {
			resMap[counterStr] = append(group, str)
		} else {
			resMap[counterStr] = []string{str}
		}
	}
	res := [][]string{}
	for _, group := range resMap {
		res = append(res, group)
	}
	return res
}
// @lc code=end

