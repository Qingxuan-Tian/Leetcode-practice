package main

import "fmt"

/**
以升序为例
首先在未排序的序列中找到最小的放到第一位，
之后重复这个过程一直往后放
*/

func selectSort(arr []int){
	for i:=0;i<len(arr)-1;i++{
		min:=i
		for j:=i+1;j<len(arr);j++{
			if arr[min]>arr[j]{
				min=j
			}
		}
		arr[i],arr[min]=arr[min],arr[i]
	}
}
func main(){
	var arr []int=[]int{1,6,4,8,9,7}
	selectSort(arr)
	fmt.Println(arr)
}
