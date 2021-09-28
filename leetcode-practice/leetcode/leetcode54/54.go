package main

import "fmt"

//offer29

func spiralOrder(matrix [][]int) []int {

	if len(matrix)==0{
		return nil
	}
	if len(matrix)==1{
		return matrix[0]
	}
	var arr []int
	if len(matrix[0])==1{
		for i:=0;i<len(matrix);i++{
			arr=append(arr, matrix[i][0])
		}
		return arr
	}

	for j:=0;j<len(matrix)-1;j++{
		i:=0
		for ;i<len(matrix[j])-j;i++{
			arr=append(arr, matrix[j][i])
		}
		i--

		m:=j+1
		if m<=len(matrix)-1-j{
			for ;m<len(matrix)-1-j;m++{
				arr= append(arr, matrix[m][i])
			}
			for ;i>=0;i--{
				arr=append(arr,matrix[m][i] )
			}
		}

	}
	return arr
}

func main()  {
	var (
		arr []int
		arr1 []int
		arr2 []int
	)
	arr=append(arr, 1)
	//arr=append(arr, 2)
	//arr=append(arr, 3)
	arr1=append(arr1, 4)
	//arr1=append(arr1, 5)
	//arr1=append(arr1, 6)
	arr2=append(arr2, 7)
	//arr2=append(arr2, 8)
	//arr2=append(arr2, 9)
	var matrix [][]int
	matrix=append(matrix, arr)
	matrix=append(matrix, arr1)
	matrix=append(matrix, arr2)
	res:= spiralOrder(matrix)
	fmt.Println(res)
}