package leetcode160
//offer52
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headB==nil||headA==nil{
		return nil
	}
	curA:= headA
	curB:= headB
	countA:=0
	countB:=0
	for curA!=curB{
		if curA==nil{
			curA=headB
			countA++
		}else {
			curA=curA.Next
		}
		if curB==nil{
			curB=headA
			countB++
		}else {
			curB=curB.Next
		}
		if countA==2||countB==2{
			return nil
		}
	}
	return curA

}












