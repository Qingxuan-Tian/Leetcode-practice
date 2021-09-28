package offer32_3

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//栈做
func levelOrder(root *TreeNode) [][]int {
	if root==nil{
		return nil
	}
	var(
		queue[]*TreeNode
		res [][]int
		count int=1
	)
	queue=append(queue, root)
	for len(queue)!=0{
		size:= len(queue)
		var res1 []int
		for i:=size-1;i>=0;i--{
			node:= queue[i]
			res1=append(res1, node.Val)
			if count%2==0{
				if node.Right!=nil{
					queue= append(queue, node.Right)
				}
				if node.Left!=nil{
					queue=append(queue,node.Left)
				}
			}else {
				if node.Left!=nil{
					queue=append(queue,node.Left)
				}
				if node.Right!=nil{
					queue=append(queue,node.Right)
				}
			}

		}
		res=append(res,res1)
		count++
		queue=queue[size:]
	}
	return res
}

func main()  {
	root := &TreeNode{
		Val: 3,
		Left: nil,
		Right: nil,
	}
	point1:= &TreeNode{
		Val: 9,
		Left: nil,
		Right: nil,
	}
	point2:= &TreeNode{
		Val: 20,
		Left: nil,
		Right: nil,
	}
	point3:= &TreeNode{
		Val: 15,
		Left: nil,
		Right: nil,
	}
	point4:= &TreeNode{
		Val: 7,
		Left: nil,
		Right: nil,
	}
	root.Left=point1
	root.Right=point2
	point2.Left=point3
	point2.Right=point4
	res:= levelOrder(root)
	fmt.Println(res)
}























