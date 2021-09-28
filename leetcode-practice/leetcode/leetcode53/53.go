package leetcode53

//offer42
//ac
//func maxSubArray(nums []int) int{
//	dp:=make([]int, len(nums))
//	dp[0]=nums[0]
//	for i:= 1;i<len(nums);i++{
//		if dp[i-1]<0{
//			dp[i]=nums[i]
//		}else {
//			dp[i]=nums[i]+dp[i-1]
//		}
//	}
//	max:= -65535
//	for _,v:= range dp{
//		if v>max{
//			max=v
//		}
//	}
//	return max
//}

//优化，dp只申请两个空间记录上一个和此时的，寻找max值也在这个过程中进行，不单独进行遍历
func maxNum(a, b int)int {
	if a>b{
		return a
	}
	return b
}
func maxSubArray(nums []int) int {
	dp:= make([]int,2)
	dp[0]=nums[0]
	max:= dp[0]
	for i:=1;i<len(nums);i++{
		if dp[(i-1)%2]<0{
			dp[i%2]=nums[i]
		}else {
			dp[i%2]=nums[i]+dp[(i-1)%2]
		}
		max=maxNum(max, dp[i%2])
	}
	return max
}




















