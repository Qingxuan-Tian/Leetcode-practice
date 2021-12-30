package main

import (
	"fmt"

	"strings"
)

func replaceSpace(s string) string {
	//========================================================
	//方法1
	//for _,v:= range s{
	//	if v==' '{
	//		res+="%20"
	//	}else {
	//		res+=string(v)
	//	}
	//
	//}
	//return res
	//===========================================================
	//方法2
	return strings.Replace(s," ","%20",-1)

}


func main(){
	var s string="hello world"
	s="abc"
	fmt.Println(s)


}











