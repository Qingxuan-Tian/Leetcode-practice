package offer10

func fib(n int) int {
	//===============================================
	//方法1
	//自顶向下+记忆化搜索
	//超时
	//memo:= make([]int,n+1)
	//if n==0{
	//	return memo[0]
	//}else if  n==1{
	//	return 1
	//}else {
	//	if memo[n]!=0{
	//		return memo[n]
	//	}else {
	//		memo[n]=fib(n-1)+fib(n-2)
	//		return memo[n]
	//	}
	//}
	//=================================================
	//方法2
	//动态规划
	if n==0{
		return 0
	}
	if n==1{
		return 1
	}
	memo:= make([]int, n+1)//0~n n+1个元素
	memo[0]=0
	memo[1]=1

	for i:=2;i<n+1;i++{
		memo[i]=(memo[i-1]+memo[i-2])%1000000007
	}
	return memo[n]





}