package main


import "fmt"

//冒泡排序
/*冒泡排序过程：
对比数组相邻两个元素的大小，若前>后，就交换两个元素。保证大的在后面
最终实现小---->大的排序
*/
 func BubbleSort(arr []int)[]int{
 	for i:=0;i<len(arr);i++{
 		for j:=0;j<len(arr)-i;j++{
 			if arr[j]>arr[j+1]{
 				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
	return arr
 }

func main()  {
	var arr []int
	arr=append(arr, 2)
	arr=append(arr, 8)
	arr=append(arr, 6)
	arr=append(arr, 1)
	arr=append(arr, 4)
	res:= BubbleSort(arr)
	fmt.Println(res)


}