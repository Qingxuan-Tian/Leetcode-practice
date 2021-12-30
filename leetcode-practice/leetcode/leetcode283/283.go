package main

import "fmt"

func moveZeroes(nums []int)  {
	var index int=0//0的个数
	i:=0
	for i<len(nums)-1-index{
		if nums[i]==0{
			for j:=i+1;j<len(nums)-index;j++{
				nums[j-1]=nums[j]
			}
			nums[len(nums)-1-index]=0
			index++
		}else{
			i++
		}
	}

}

func main(){
	var nums []int=[]int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}
