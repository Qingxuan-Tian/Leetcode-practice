package leetcode79
func findFirstPointIndex(board [][]byte, firstByte byte)[][]int{
		var res [][]int
		for i:=0;i<len(board);i++{
			for j:=0;j<len(board[0]);j++{
				if board[i][j]==firstByte{
					res=append(res, []int{i,j})
				}
			}
		}
		return res
}
func exist(board [][]byte, word string) bool {
	var (
		width int=len(board[0])
		length int=len(board)
	)


	firstIndex:=findFirstPointIndex(board,byte(word[0]))
	for _,index:= range firstIndex{
		//标志位矩阵初始化
		symbol:= make([][]int, length)
		for i:=0;i<length;i++{
			symbol[i]= make([]int,width)
		}
		i:= index[0]
		j:= index[1]
		symbol[i][j]=1
		for i<len(board)&&i>=0&&j<len(board[0])&&j>=0{

		}

	}
}

