/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	for _, c := range root.Children {
		res = append(res, preorder(c)...)
	}
	return res
}

func preorderStack(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append(res, root.Val)
		for i := len(root.Children) - 1; i >= 0; i-- {
			stack = append(stack, root.Children[i])
		}
	}
	return res
}
// @lc code=end

