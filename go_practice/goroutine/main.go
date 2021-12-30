package main

import (
	"fmt"
	"time"
)

//协程相关
func main(){
	//var a [10]int
	for i:=0;i<10;i++{
		go func() {
			for {
				fmt.Println(i)
			}
		}()
	}
	//for i:=0;i<9;i++{
	//	go func() {
	//		a[i]++
	//	}()
	//}
	time.Sleep(time.Millisecond)
	//fmt.Println(a)

}


