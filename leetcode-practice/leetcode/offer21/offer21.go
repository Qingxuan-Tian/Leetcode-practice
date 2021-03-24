package offer21
//滑动窗口
func exchange(nums []int) []int {
	size:= len(nums)
	index, cur:= 0,1
	for index<size{
		if nums[cur]%2!=0{
			//为奇数
			cur++
			if cur==index{
				index++
			}
		}else {
			//为偶数
			nums[cur],nums[index]=nums[index],nums[cur]
			index++
		}
	}
	return nums
}