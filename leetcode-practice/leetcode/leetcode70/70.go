package leetcode70

func numWays(n int) int {
	//========================================================
	//递归，自顶向下+记忆化搜索
	//超时
	//if n==0||n==1{
	//	return 1
	//}
	//memo:= make([]int,n+1)
	//if memo[n]==0{
	//	memo[n]=(numWays(n-1)+numWays(n-2))%1000000007
	//}
	//return memo[n]
	//==========================================================
	//方法2：动态规划
	memo:= make([]int, n+1)
	if n==1||n==0{
		return 1
	}
	memo[0]=1
	memo[1]=1
	for i:=2;i<=n;i++{
		memo[i]=memo[i-1]+memo[i-2]
	}
	return memo[n]
}
