package main

import "fmt"

//方法1：快排
//========================================================================================
func partition(arr []int, left int, right int) int {

	index:= left-1
	target:= arr[right]
	for j:=left;j<right;j++{
		if arr[j]<= target{
			index++
			arr[index],arr[j]=arr[j],arr[index]
		}
	}
	arr[index+1],arr[right]=arr[right],arr[index+1]
	return index+1
}

func _quickSort(arr []int, left int, right int)  {
	if left<right{
		index:= partition(arr, left,right)
		_quickSort(arr, left, index-1)
		_quickSort(arr, index+1,right)

	}
}

func quickSort(arr []int) []int{
	index:= partition(arr, 0, len(arr)-1)
	_quickSort(arr,0,index-1)
	_quickSort(arr,index+1, len(arr)-1)
	return arr
}
func getLeastNumbers(arr []int, k int) []int {
	arr=quickSort(arr)
	if k==0{
		return nil
	}
	res:= arr[0:k]
	return res
}
//=================================================================================
//方法2：建堆，大根堆
func buildHeap(arr []int,start int, end int){
	root:= start
	child:= 2*root+1
	if child>end{
		return
	}
	//看左右两个子节点谁大，大的和root比
	if child+1<=end&&arr[child]<=arr[child+1]{
		child++
	}
	if arr[child]>arr[root]{
		arr[child],arr[root]=arr[root],arr[child]
		root=child
		buildHeap(arr,root,end)
	}else {
		return
	}
	return

}
//把数组堆化
func init_arr(arr []int)[]int{
	first := len(arr)/2-1
	for i:= first;i>-1;i--{
		buildHeap(arr, i, len(arr)-1)
	}
	return arr
}
//利用堆结构排序
func heap_Sort(arr []int)[]int{
	var (
		res []int
	)
	arr1:= arr
	for {
		if len(arr1)==0{
			break
		}
		res=append(res, arr1[0])
		arr1[0],arr1[len(arr1)-1]=arr1[len(arr1)-1],arr1[0]
		arr1=arr1[:len(arr1)-1]
		buildHeap(arr1,0, len(arr1)-1)
	}
	return res
}
func getLeastNumbers1(arr []int, k int) []int{
	if k==0{
		return nil
	}
	arr=init_arr(arr)
	arr=heap_Sort(arr)
	res:=arr[len(arr)-k:]
	return res
}
//===============================================================================================================

func main(){
	var arr []int
	arr=append(arr, 0)
	arr=append(arr, 1)
	arr=append(arr, 2)
	arr=append(arr, 1)

	arr=init_arr(arr)
	fmt.Println(arr)

	res:= getLeastNumbers1(arr,1)
	fmt.Println(arr)
	fmt.Println(res)

}