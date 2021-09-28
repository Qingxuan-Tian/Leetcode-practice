package main

import (
	"fmt"
	"time"
)

//import (
//	"./leetcode70"
//	"fmt"
//)

//绝对路径
//import "../leetcode/linklist"
//相对路径
//import "./linklist"
var m= make(map[int]int)
func test(arr []int){
	for _,v:= range arr{
		fmt.Println(v)
	}
}
func store(key, val int)  {
	m[key]=val
	fmt.Println("store key ",key)
}
func load(key int) int {
	fmt.Println("load key", key)
	return m[key]
}
func main (){
	for i:=0;i<3;i++{
		go store(i,i)
		go load(i)
	}
	time.Sleep(time.Millisecond*100)
	fmt.Println(m)

}
