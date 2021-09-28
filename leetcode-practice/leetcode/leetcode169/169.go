package leetcode169
//offer39

import (
	"fmt"
	"math"
)
//通过
func majorityElement(nums []int) int {
	//target:=math.Floor(float64(len(nums))/2+0.5)
	var target int
	len:= len(nums)
	if len%2==0{
		target=len/2
	}else {
		target=len/2+1
	}
	m:= make(map[int]int)
	for _,v:= range nums{
		m[v]++
		if n,_:=m[v];n>=int(target){
			return v
		}
	}
	return -1
}
//摩尔投票法
func majorityElement1(nums []int) int{
	var(
		votes int=0
		x int
	)
	for _,v:= range nums{
		if votes==0{
			x=v
		}
		if v==x{
			votes++
		}else {
			votes--
		}
	}
	return x
}

func main()  {
	var a int=5
	b:= a/2
	c:= float64(a)/2
	d:= math.Ceil(c)
	fmt.Println(a,b,c,d)
}