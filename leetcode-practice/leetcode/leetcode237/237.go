package leetcode237

/**
ac
因为只知道当前节点，不能获取到节点的所有信息，所以，采用改变节点值的方式
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
func deleteNode(node *ListNode) {
	if node==nil{
		return
	}
	if node.Next==nil{
		node=nil
		return
	}
	node.Val=node.Next.Val
	node.Next=node.Next.Next
	return
}