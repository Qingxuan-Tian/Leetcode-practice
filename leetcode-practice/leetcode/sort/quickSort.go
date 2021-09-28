package main

import "fmt"

func partition(arr []int, left int, right int) int {
	if left>=right{
		return left
	}
	target:= arr[left]
	index:= left+1
	for i:=left+1;i<=right;i++{
		if arr[i]<target{
			arr[i],arr[index]=arr[index],arr[i]
			index++
		}
	}
	arr[left],arr[index-1]=arr[index-1],arr[left]
	return index-1
}
func quickSort(arr []int, left int,right int){
	if left<right{
		index:= partition(arr,left,right)
		quickSort(arr,left,index-1)
		quickSort(arr,index+1,right)
	}


}
func main(){
	var arr []int=[]int{1,6,4,8,7,9}
	quickSort(arr, 0,len(arr)-1)
	fmt.Println(arr)

}
