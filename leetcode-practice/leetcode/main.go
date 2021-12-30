package main

import "fmt"

//func findSum(arr []int, target int)[][]int{
//	if len(arr)==0{
//		return [][]int{}
//	}
//	var res [][]int
//	sort.Slice(arr, func(i, j int) bool {
//		return arr[i]<arr[j]
//	})
//	for count:=0;count<len(arr);count++{
//		r:= findSum(arr,target-arr[count])
//		for i:=0;i<len(r);i++{
//			r[i]=append(r[i],arr[count])
//		}
//		res=append(res,r ... )
//	}
//
//	return res
//
//}
func main(){
	//var arr [][]int
	//var a []int
	//fmt.Println(len(arr))
	//arr=append(arr,a)
	//fmt.Println(len(arr))
	//fmt.Println(arr)
	//fmt.Println(a)

	var arr1 []int=[]int{1,2,3}
	var arr2 []int=[]int{7,8}
	fmt.Printf("arr1:%p, arr2:%p\n",&arr1,&arr2)
	copy(arr2,arr1)
	fmt.Printf("arr1:%p, arr2:%p\n",&arr1,&arr2)
	fmt.Println(arr1,arr2)//[1,2,3]  [1,2]
}
