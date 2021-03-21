package offer5

import "strings"

func replaceSpace(s string) string {
	//========================================================
	//方法1
	//var res string
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