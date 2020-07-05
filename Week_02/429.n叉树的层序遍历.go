/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	traverse := func(node *Node, level int) {}
	traverse = func(node *Node, level int) {
		if len(res) == level {
			res = append(res, []int{node.Val})
		} else {
			res[level] = append(res[level], node.Val)
		}
		for _, c := range node.Children {
			traverse(c, level + 1)
		}
	}
	traverse(root, 0)
	return res
}
// @lc code=end

