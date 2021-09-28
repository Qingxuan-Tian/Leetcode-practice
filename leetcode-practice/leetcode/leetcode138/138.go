package main

import "fmt"

// Definition for a Node.
 type Node struct {
     Val int
     Next *Node
    // Random *Node
 }


//func copyRandomList(head *Node) *Node {
//	var headCopy *Node
//	cur:= head
//	for cur!=nil{
//
//	}
//}

func main(){
	for i:=1;i<2;i++{
		point := new(Node)
		point1:=new([]int)
		fmt.Println(point)
		fmt.Println(*point1)
	}
}