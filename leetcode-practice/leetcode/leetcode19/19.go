package leetcode19
// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead:= &ListNode{
		Next: head,
	}
	//双指针初始化
	p:= dummyHead
	q:= dummyHead
	//指针q初始化
	for i:=0;i<n+1;i++{
		q=q.Next
	}
	//开始找需要删除的节点的前一个节点位置，即为p
	for q!=nil{
		p=p.Next
		q=q.Next
	}
	//已经找到待删除节点的前一个节点，进行删除操作
	p.Next=p.Next.Next
	return dummyHead.Next


}
