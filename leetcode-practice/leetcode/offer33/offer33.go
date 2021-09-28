package offer33

import "fmt"

func verifyPostorder(postorder []int) bool {
	if len(postorder)<=1{
		return true
	}
	root:= postorder[len(postorder)-1]
	index:= 0
	for k,v := range postorder{
		//这里是大于等于，因为如果发生右子树为空的情况那inedx就是左子树最后一个，也就是root前一个元素
		if v>=root{
			index=k
			break
		}
	}
	leftTree:= postorder[:index]
	rightTree:= postorder[index:len(postorder)-1]
	for _,v:= range leftTree{
		if v>root{
			return false
		}
	}
	for _, v:= range rightTree{
		if v<root{
			return false
		}
	}
	//还要再检查一下左右子树自己是否满足，因为也许左子树都<root但是内部不符合
	return verifyPostorder(leftTree)&&verifyPostorder(rightTree)
}

func main(){
	var arr []int=[]int{4,6,7,5}
	res:= verifyPostorder(arr)
	fmt.Println(res)
}
















