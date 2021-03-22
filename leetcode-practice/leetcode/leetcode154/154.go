package leetcode154
func minArray(numbers []int) int {
	//=============================================
	//方法1，不好，没有利用到这是一个翻转数组的条件
	//min:=65535
	//for i:=0;i<len(numbers);i++{
	//	if numbers[i]<=min{
	//		min=numbers[i]
	//	}
	//}
	//return min
	//===================================================
	//方法2：二分搜索法
	left:=0
	right:= len(numbers)-1
	for left<right{
		mid:= (left+right)/2
		if numbers[mid]>numbers[right]{
			left=mid+1
		}else {
			right=mid
		}

	}
	return numbers[right]
}
