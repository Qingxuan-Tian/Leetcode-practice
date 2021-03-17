package linklist

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}


//生成测试链表
func CreatLinkList(arr [2]int, n int)*ListNode{
	var dummyHead *ListNode=new(ListNode)
	curPoint := dummyHead

	for _,v:= range arr{
		point:= new(ListNode)
		point.Next=nil
		point.Val=v
		curPoint.Next=point
		curPoint=curPoint.Next
	}
	return dummyHead.Next
}
//打印测试链表
func PrintLinkedList(head *ListNode){
	curPoint := head
	for curPoint!=nil {
		fmt.Println(curPoint.Val,"->")
		curPoint=curPoint.Next
	}
	fmt.Println("nil")
}

//获取链表长度
func getLinkListSize(head *ListNode)int  {
	num:=0
	cur:= head
	for cur!=nil{
		num++
		cur=cur.Next
	}
	return num

}