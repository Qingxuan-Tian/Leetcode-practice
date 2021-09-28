package main

import (
	"fmt"
	"io"
)

func mySqrt(x int) int {
	for i := 0; i <= x; i++ {
		res := i * i
		if res == x {
			return i
		} else if res > x {
			return i - 1
		}
	}
	return -1
}

func finSum(num int, m int) int{
	var res int=0
	memo:= make([]int, m)
	memo[0]=num
	for i:=1;i<m;i++{
		memo[i]=mySqrt(memo[i-1])
		res+= memo[i]
	}
	res=res+num
	return res
}

func main()  {
	var (
		num int
		m int
	)
	for {
		_,err:= fmt.Scan(&num,&m)
		if err==io.EOF{
			break
		}
		fmt.Println(finSum(num,m))
	}

}