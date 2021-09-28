package main
import "fmt"

//冒泡排序
/*冒泡排序过程：
对比数组相邻两个元素的大小，若前>后，就交换两个元素。保证大的在后面
最终实现小---->大的排序
注意i和j的范围，i从0到len(arr）-1；j从1到len（arr）-i-1；每次j-1位置和j位置进行比较
*/

func BubbleSort(arr []int){
	if len(arr)==0{
		return
	}
	for i:=0;i<len(arr);i++{
		for j:=1;j<len(arr)-i;j++{
			if arr[j-1]>arr[j]{
				arr[j-1],arr[j]=arr[j],arr[j-1]
			}
		}
	}

}

func main()  {
	var arr []int=[]int{1,-1,0}
	BubbleSort(arr)
	fmt.Println(arr)


}