package offer57

//忽略数组递增条件,ac 时间复杂度和空间复杂度不好
func twoSum1(nums []int, target int) []int {

	m:= make(map[int]int)
	for _,v:= range nums{
		if _,ok:= m[target-v];ok{
			return []int{v,target-v}
		}else {
			m[v]++
		}
	}
	return nil
}
//考虑数组递增条件,二分法找，ac，时间空间复杂度不好
func findNumIndex(nums []int, left, right int, number int)int{
	//left, right 闭区间
	for left<= right{
		mid:=(left+right)/2
		if nums[mid]<number{
			left=mid+1
			continue
		}else if nums[mid]>number{
			right=mid-1
			continue
		}else {
			return mid
		}
	}
	return -1
}
func twoSum2(nums []int,target int)[]int{
	for k,v:= range nums{
		if target-v>v{
			if index:= findNumIndex(nums,k+1, len(nums)-1,target-v);index!=-1{
				return []int{v, nums[index]}
			}
		}
	}
	return nil
}
//方法3,双指针，O(1)复杂度
func twoSum(nums []int, target int)[]int{
	var(
		left int=0
		right int=len(nums)-1
	)
	//先过滤一些绝对不可能的选项，加了这一步时间复杂度减少！
	for nums[right]>target{
		right--
	}
	for left<right{
		temp := nums[left]+nums[right]
		if temp<target{
			left++
		}else if temp>target{
			right--
		}else {
			return []int{nums[left], nums[right]}
		}
	}
	return nil
}