package leetcode21

import "../linklist"
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getSize(l *linklist.ListNode) int{
	cur:= l
	size :=0
	for cur!=nil{
		size++
		cur=cur.Next
	}
	return size
}
func mergeTwoLists(l1 *linklist.ListNode, l2 *linklist.ListNode) *linklist.ListNode {
	//===========================================================================================================
	//情况考虑的不全
	//size1:= getSize(l1)
	//size2:= getSize(l2)
	//if size1==0&&size2!=0{
	//	return l2
	//}else if size2==0&&size1!=0{
	//	return l1
	//}else if size1==0&&size2==0{
	//	return nil
	//}
	//cur1:= l1
	//
	//cur2:= l2
	//
	//temp := &linklist.ListNode{
	//	Next: l1,
	//}
	//head:= temp
	//for cur1!=nil&&cur2!=nil{
	//	next1:= cur1.Next
	//	next2:= cur2.Next
	//	if cur1.Val<= cur2.Val{
	//		temp.Next=cur1
	//		cur1.Next=cur2
	//		temp=cur2
	//	}else {
	//		temp.Next=cur2
	//		cur2.Next=cur1
	//		temp=cur1
	//	}
	//	cur1=next1
	//	cur2=next2
	//}
	//if cur1==nil{
	//	if cur2!=nil{
	//		temp.Next=cur2
	//	}
	//}else {
	//	temp.Next=cur1
	//}
	//return head.Next
//=============================================================================================================
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}
	head := &linklist.ListNode{}
	if l1.Val>l2.Val{
		head= l2
		head.Next=mergeTwoLists(l1,l2.Next)
	}else {
		head=l1
		head.Next=mergeTwoLists(l1.Next,l2)
	}
	return head

}















