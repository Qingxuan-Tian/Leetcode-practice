package main

import (
	"fmt"
)

func cal(x float64, n int) float64 {
	for n!=0{
		x=x*x
		n--
	}
	return x
}
//递归：超时
//func myPow(x float64, n int) float64 {
//	if n==0{
//		return 1.00000
//	}
//	if n==1{
//		return x
//	}
//	var res float64
//	if n>0{
//		if n%2==0{
//			res=myPow(x, n/2)*myPow(x,n/2)
//		}else {
//			res=x*myPow(x,(n-1)/2)*myPow(x,(n-1)/2)
//		}
//
//	}else {
//		n=int(math.Abs(float64(n)))
//		x=1/x
//		if n%2==0{
//			res=myPow(x, n/2)*myPow(x,n/2)
//		}else {
//			res=x*myPow(x,(n-1)/2)*myPow(x,(n-1)/2)
//		}
//	}
//	return res
//}

//n为最大时超内存限制
//func myPow(x float64, n int)float64{
//	if n==0{
//		return 1
//	}
//	if n==1{
//		return x
//	}
//	if n==-1{
//		return 1/x
//	}
//
//	if n<0{
//		//n=int(math.Abs(float64(n)))
//		n=-n
//		x=1/x
//	}
//	memo:= make([]float64, n+1)
//	memo[0]=1
//	memo[1]=x
//========================================================
//内存没超，超时
//	//memo:= make([]float64, 2)
//	//memo[1]=x
//	//for i:=2;i<n+1;i++{
//	//	memo[i%2]=memo[(i-1)%2]*x
//	//}
//	//if n%2==0{
//	//	return memo[0]
//	//}else {
//	//	return memo[1]
//	//}
//==================================================================
//
//	for i:=2;i<n+1;i++{
//		if i%2==0{
//			memo[i]=memo[i/2]*memo[i/2]
//		}else {
//			memo[i]=x*memo[(i-1)/2]*memo[(i-1)/2]
//		}
//	}
//
//	return memo[n]
//}
//正确！！
func myPow(x float64, n int)float64{
	if n==0{
		return 1
	}
	if n==1{
		return x
	}
	if n<0{
		n=-n
		x=1/x
	}
	temp:= myPow(x, n/2)
	if n%2==0{
		return temp*temp
	}else {
		return x*temp*temp
	}
}

func main() {
	var i int =-1
	k:= ^i+1
	m:= i<<1
	j:= i&0xffff
	fmt.Printf("%x",i)
	fmt.Println(i,j,k,m)
}