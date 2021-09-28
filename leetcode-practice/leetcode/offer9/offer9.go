package offer9
//ac
import (
	"container/list"
	"fmt"
)

//用切片实现队列
//type CQueue struct {
//	arr []int
//	tail int
//}
//
//func Constructor() CQueue {
//	queue := CQueue{
//		arr: []int{},
//		tail: 0,
//	}
//	return queue
//}
//
//
//func (this *CQueue) AppendTail(value int)  {
//	this.arr=append(this.arr, value)
//	this.tail++
//}
//
//
//func (this *CQueue) DeleteHead() int {
//	if len(this.arr)==0{
//		//panic("wrong")
//		//题目要求返回-1
//		return -1
//	}
//	//if this.tail==0{
//	//	panic("wrong")
//	//}
//	temp:= this.arr[0]
//	this.arr=this.arr[1:]
//	this.tail--
//	return temp
//}

//栈实现队列。双栈
//ac
type CQueue struct {
	stackin *list.List
	stackout *list.List
}
func Constructor()CQueue{
	return CQueue{
		stackin: list.New(),
		stackout: list.New(),
	}
}
func (this *CQueue)AppendTail(value int){
	this.stackin.PushBack(value)
}

func (this *CQueue)DeleteHead()int{
	if this.stackout.Len()==0{
		for this.stackin.Len()>0{
			this.stackout.PushBack(this.stackin.Remove(this.stackin.Back()))
		}
	}
	if this.stackout.Len()>0{
		e:= this.stackout.Back()
		this.stackout.Remove(e)
		return e.Value.(int)
	}
	return -1
}

func main()  {
	//var arr1 []int
	//if arr1==nil{
	//	fmt.Println("ok")
	//}else {
	//	fmt.Println("no")
	//}
	//arr2:= make([]int,0)
	//if arr2==nil{
	//	fmt.Println("ok")
	//}else {
	//	fmt.Println("no")
	//}
	//fmt.Println(arr1[0])
	//fmt.Println(arr2)

	 obj := Constructor()
	obj.AppendTail(3)
	 //obj.AppendTail(1);
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())


}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */