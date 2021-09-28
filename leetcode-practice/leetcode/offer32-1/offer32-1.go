//package offer32_1
package offer32_1

import (
	"fmt"
	"time"
)

// Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

//队列做
//ac
func levelOrder(root *TreeNode) []int {
	if root==nil{
		return nil
	}
	var (
		queue []*TreeNode
		res []int
	)
	queue=append(queue, root)
	for len(queue)!=0{
		res=append(res, queue[0].Val)
		if queue[0].Left!=nil{
			queue=append(queue, queue[0].Left)
		}
		if queue[0].Right!=nil{
			queue= append(queue, queue[0].Right)
		}
		queue=queue[1:]
	}
	return res

}



//func levelOrder(root *TreeNode) []int {
//	var res []int
//	if root==nil{
//		return res
//	}
//	queue := []*TreeNode{root}
//	for len(queue)>0{
//		size :=len(queue)
//		for i:=0;i<size;i++{
//			node :=queue[i]
//			if node.Left!=nil{
//				queue = append(queue,node.Left)
//			}
//			if node.Right!=nil{
//				queue = append(queue,node.Right)
//			}
//			res = append(res,node.Val)
//		}
//		queue = queue[size:]
//	}
//	return res
//}



func main(){
	var arr []int=[]int{1,2,3,4,5,6}
	t:= time.Now()
	fmt.Println("now")
	fmt.Println(t)
	arr=arr[3:]
	//arr=arr[1:]
	//arr=arr[1:]
	//arr=arr[1:]
	fmt.Println("after")
	fmt.Println(t.Sub(time.Now()))
}















