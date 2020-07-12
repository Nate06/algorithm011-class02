/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	tree := &TreeNode{Val: preorder[0]}

	pos := 0
	for k, v := range inorder {
		if v == preorder[0] {
			pos = k
			break
		}
	}
	tree.Left = buildTree(preorder[1:len(inorder[:pos]) + 1], inorder[:pos])
	tree.Right = buildTree(preorder[len(inorder[:pos]) + 1:], inorder[pos + 1:])
	return tree
}
// @lc code=end

