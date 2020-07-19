/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 */

// @lc code=start
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordMap := map[string]int{}
	for i, word := range wordList {
		wordMap[word] = i
	}

	if _, ok := wordMap[beginWord]; !ok {
		wordList = append(wordList, beginWord)
		wordMap[beginWord] = len(wordList) - 1
	}

	edges := make([][]int, len(wordList))
	for i, word := range wordList {
		for j, diffWord := range wordList {
			if isDiffOne(word, diffWord) {
				edges[i] = append(edges[i], j)
			}
		}
	}

	cost := make([]int, len(wordList))
	for i := range wordList {
		cost[i] = math.MaxInt32
	}
	cost[wordMap[beginWord]] = 0

	queue := [][]int{{wordMap[beginWord]}}
	res := [][]string{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		last := current[len(current) - 1]

		if wordList[last] == endWord {
			list := []string{}
			for _, i := range current {
				list = append(list, wordList[i])
			}
			res = append(res, list)
		}

		for _, edge := range edges[last] {
			if cost[edge] >= cost[last] + 1 {
				cost[edge] = cost[last] + 1
				temp := make([]int, len(current))
				copy(temp, current)
				temp = append(temp, edge)
				queue = append(queue, temp)
			}
		}
	}
	return res
}

func isDiffOne(word1, word2 string) bool {
	diff := 0
	for i := range word1 {
		if word1[i] != word2[i] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}
	return diff == 1
}
// @lc code=end

