package main

import (
	"fmt"
)

//二分查找法
func search(nums []int, target int) int {
	var(
		left int=0
		right int=len(nums)-1
		mid int=0
		lIndex int
		rIndex int
	)
	//find left index
	for left<=right{
		mid=(left+right)/2
		if nums[mid]==target{
			if left==right{
				lIndex=mid
				break
			}else{
				right=mid
			}
		}else if nums[mid]<target{
			left=mid+1
		}else{
			right=mid-1
		}
	}
	//test exit or not
	if left>right{
		return 0
	}

	//find right index
	left=lIndex+1
	right=len(nums)-1
	for left<=right{
		mid=(left+right)/2
		if nums[mid]==target{
			if left==right{
				rIndex=mid
				break
			}
			right=mid
		}else if nums[mid]<target{
			left=mid+1
		}else{
			right=mid-1
		}
	}
	if rIndex<lIndex{
		rIndex=lIndex
	}
	return (rIndex-lIndex)+1
}

func main(){
	var nums []int=[]int{1,4}
	fmt.Println(search(nums,4))
}







