package main

import (
	"fmt"
)

//offer 53-1
//方法一：map做 ac复杂度O（n）,没有用上关键信息有序数组
//func search(nums []int, target int) int {
//	//map做
//	m:= make(map[int]int)
//	for _,v:= range nums{
//		m[v]++
//	}
//	if val,ok:= m[target];ok{
//		return val
//	}else {
//		return 0
//	}
//
//}

//方法二
func search(nums []int, target int)int{
	//边界情况，为空
	if len(nums)==0{
		return 0
	}
	count:=0
	left:=0
	right:= len(nums)
	var mid int=0
	for left<right{
		mid= (left+right)/2
		if nums[mid]<target{
			left=mid+1
		}else {
			right=mid
		}
	}
	lower_index:= left
	right=len(nums)
	for left<right{
		mid=(left+right)/2
		if nums[mid]<=target{
			left=mid+1
		}else {
			right=mid
		}
	}
	upper_index:= left
	count=upper_index-lower_index

	return count

}

func main()  {
	arr:= []int{1,2,3,3,3,3,4,5,9}
	res:=search(arr,10)
	fmt.Println(res)
}