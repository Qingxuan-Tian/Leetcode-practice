package leetcode101
//offer28
import (
	"fmt"
)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//镜像一个二叉树,注意这里！！！！不可以直接对传进来的二叉树指针进行操作，这样原来的二叉树也发生改变
func mirrorTree(root *TreeNode) *TreeNode {

	if root ==nil{
		return root
	}
	root1:= *root
	root1.Left=mirrorTree(root1.Left)
	root1.Right=mirrorTree(root1.Right)
	root1.Right, root1.Left= root1.Left, root1.Right
	return &root1
}

func isSame(root *TreeNode, root1 *TreeNode)bool  {

	if root1==nil&&root==nil{
		return true
	}
	if root1==nil||root==nil{
		return false
	}
	var ok bool

	if root.Val==root1.Val{
		ok= isSame(root.Right,root1.Right) && isSame(root.Left, root1.Left)
	}else {
		return false
	}
	return ok
}

func isSymmetric(root *TreeNode) bool {
	if root==nil{
		return true
	}
	root1:= mirrorTree(root)
	res:= isSame(root, root1)
	return res
}

func main()  {
	root:= TreeNode{
		Val: 1,
	}
	point:=TreeNode{
		Val: 2,
		Left: nil,
		Right: nil,
	}
	root.Left=&point
	root.Right=nil
	res:= isSymmetric(&root)
	fmt.Println(res)

}