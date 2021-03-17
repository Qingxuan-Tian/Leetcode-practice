package leetcode206

//Definition for singly-linked list.
type ListNode struct {
  Val int
  Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var(
		pre *ListNode=nil
		cur *ListNode=head
		next *ListNode
	)
	for cur!=nil{
		next=cur.Next//确保在cur不为空的情况下访问cur的next
		cur.Next=pre
		pre=cur
		cur=next
	}
	return pre
}

