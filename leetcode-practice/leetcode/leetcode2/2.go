package leetcode2
//ac
import "fmt"

// Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head:= &ListNode{}
	cur:= head
	var (
		val int=0
		carry int=0
		temp int=0
		//carryPre int=0
	)
	for l1!=nil||l2!=nil{
		if l1==nil||l2==nil{
			if l1==nil{
				temp=l2.Val+carry
			}else {
				temp=l1.Val+carry
			}
		}else {
			temp=l1.Val+l2.Val+carry

		}
		if temp>=10{
			carry= temp/10
			temp=temp%10
		}else {
			carry=0
		}
		val=temp
		cur.Next=&ListNode{
			Val: val,
			Next: nil,
		}
		cur=cur.Next
		if l1!=nil{
			l1=l1.Next
		}
		if l2!=nil{
			l2=l2.Next
		}

	}
	if carry!=0{
		point:= &ListNode{
			Val: carry,
			Next: nil,
		}
		cur.Next=point
	}
	return head.Next

}

func main()  {
	point1:=&ListNode{
		Val: 9,
		Next: nil,
	}
	point2:=&ListNode{
		Val: 9,
		Next: nil,
	}
	point3:=&ListNode{
		Val: 9,
		Next: nil,
	}
	point4:=&ListNode{
		Val: 9,
		Next: nil,
	}
	point5:=&ListNode{
		Val: 9,
		Next: nil,
	}
	//point6:=&ListNode{
	//	Val: 9,
	//	Next: nil,
	//}
	l1:= point1
	l1.Next=point2
	point2.Next=point3
	l2:= point4
	l2.Next=point5
	//point5.Next=point6
	res:= addTwoNumbers(l1,l2)
	for res!=nil{
		fmt.Println(res.Val)
		res=res.Next
	}

}

