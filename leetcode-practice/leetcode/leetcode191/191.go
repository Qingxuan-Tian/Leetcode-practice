package leetcode191

func hammingWeight(num uint32) int {
	if num==0{
		return 0
	}
	res:=0
	//模式1，双百
	//for num!=0{
	//	if num&1==1{
	//		res++
	//
	//	}
	// num=num>>1
	//}
	//模式2
	for ; num!=0;num>>=1{
		if num&1==1{
			res++
		}
	}
	return res
}

