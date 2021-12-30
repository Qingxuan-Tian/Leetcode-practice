package main

import "fmt"

/**
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。

*/
//ac
type MinStack struct {
	stack []int
	top int
}


func Constructor() MinStack {
	var Mstack MinStack
	Mstack.stack=make([]int,0)
	Mstack.top=0
	return Mstack
}


func (this *MinStack) Push(val int)  {
	this.stack=append(this.stack, val)
	this.top++
}


func (this *MinStack) Pop()  {
	if len(this.stack)==0{
		panic("the stack is empty")
		return
	}
	this.stack=this.stack[:len(this.stack)-1]
	this.top--
}


func (this *MinStack) Top() int {
	if len(this.stack)==0{
		panic("the stack is empty")
		return -1
	}
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	if len(this.stack)==0{
		panic("the stack is empty")
		return -1
	}
 	var min int = int(^uint(0)>>1)
	for _,v:=range this.stack{
		if v<min{
			min=v
		}
	}
	return min
}


func main(){
	obj:= Constructor()
	obj.Push(1)
	obj.Push(10)
	obj.Push(-1)
	res:=obj.GetMin()
	fmt.Println(res)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
