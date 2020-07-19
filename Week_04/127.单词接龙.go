/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 0
	}
	visited := map[string]bool{}
	queue := []string{beginWord}
	level := 1
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			current := queue[i]
			if current == endWord {
				return level
			}
			for _, word := range wordList {
				diff := 0
				for j := range current {
					if current[j] != word[j] {
						diff++
					}
				}
				if diff == 1 && !visited[word] {
					visited[word] = true
					queue = append(queue, word)
				}
			}
		}
		queue = queue[size:]
		level++
	}
	return 0
}
// @lc code=end

