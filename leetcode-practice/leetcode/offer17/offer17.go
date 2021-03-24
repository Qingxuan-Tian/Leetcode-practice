package offer17

import (
	"math"
)

func printNumbers(n int) []int {
	max:= 0
	for n!=0{
		temp:= math.Pow(10,float64(n-1))

		max+=9*int(temp)
		n--
	}
	res:= make([]int, max)
	for max!=0{
		res[max-1]=max
		max--
	}
	return res
}

