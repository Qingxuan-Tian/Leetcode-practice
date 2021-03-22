package offer6

import "../linklist"
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *linklist.ListNode) []int {
	var (
		res []int
		temp int
	)
	cur:= head
	for cur!=nil{
		res= append(res, cur.Val)
		cur= cur.Next
	}
	size := len(res)
	if res==nil||size==0{
		return nil
	}
//------------------------------------------------------------
//方法1
	l,r := 0,size-1
	if size%2==0{
		for l!=size/2{
			temp=res[l]
			res[l]=res[r]
			res[r]=temp
			l++
			r--
		}
	}else {
		for l!=r{
			temp=res[l]
			res[l]=res[r]
			res[r]=temp
			l++
			r--
		}
	}
	//---------------------------------------------------
	//方法2
	//for i, j := 0, len(res)-1; i < j; {
	//	res[i], res[j] = res[j], res[i]
	//	i++
	//	j--
	//}
//---------------------------------------------------------
	return res


}