package leetcode226
//offer27
import "../binaryTree"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *binaryTree.TreeNode) *binaryTree.TreeNode {
	if root==nil{
		return root
	}
	root.Left=mirrorTree(root.Left)
	root.Right=mirrorTree(root.Right)
	root.Right,root.Left=root.Left,root.Right
	return root
}