package leetcode105

import "../binaryTree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *binaryTree.TreeNode {
	if len(preorder)==0{
		return nil
	}
	root:= &binaryTree.TreeNode{
		Val: preorder[0],
		Left: nil,
		Right: nil,
	}
	//从中序遍历中找到root位置
	index:=0
	for k,v:=range inorder{
		if v==preorder[0]{
			index=k
		}
	}
	root.Left=buildTree(preorder[1:len(inorder[:index])+1],inorder[:index])
	root.Right=buildTree(preorder[1+len(inorder[:index]):],inorder[index+1:])
	return root
}








