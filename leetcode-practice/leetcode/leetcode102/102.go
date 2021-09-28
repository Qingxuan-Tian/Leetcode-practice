package leetcode102
//offer32_2

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//ac
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root==nil{
		return nil
	}
	var (
		queue []*TreeNode
		res [][]int
	)
	queue= append(queue, root)
	for len(queue)!=0{
		size := len(queue)
		var res1 []int
		for i:=0;i<size;i++{
			res1=append(res1, queue[i].Val)
			if queue[i].Left!=nil{
				queue=append(queue,queue[i].Left )
			}
			if queue[i].Right!=nil{
				queue=append(queue,queue[i].Right)
			}
		}
		res= append(res, res1)
		queue=queue[size:]
	}
	return res
}