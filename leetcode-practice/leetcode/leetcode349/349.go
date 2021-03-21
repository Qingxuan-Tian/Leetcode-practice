package leetcode349

func intersection(nums1 []int, nums2 []int) []int {
	type void struct {}
	var (
		member void
		arr []int
	)
	set1 := make(map[int]void)
	//遍历一遍num1
	for _,v:= range nums1{
		//这一步判断多余了
		if _,ok:= set1[v];!ok{
			set1[v]=member
		}

	}
	//遍历num2找和num1相同元素
	for _,v:=range nums2{
		if _,ok:= set1[v];ok{
			delete(set1,v)
			arr=append(arr, v)
		}
	}
	return arr
}