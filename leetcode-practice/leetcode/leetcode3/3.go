package leetcode3

//import "strings"
//
//func lengthOfLongestSubstring(s string) int {
//	var l,r,res=0,-1,0
//	var freq[256] int
//	var len = strings.Count(s,"")-1
//	for {
//		if r+1<len && freq[s[r+1]]==0{
//			r++
//			freq[s[r]]++
//		}else {
//			if(len==0){//考虑特殊情况，数组长度为0
//				res=0
//			}else{
//				freq[s[l]]--
//				l++
//			}
//		}
//		if res<r-l+1{
//			res=r-l+1
//		}
//		if l>=len{
//			break
//		}
//	}
//	return res
//}
func maxNum(a,b int)int{
	if a>b{
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	var(
		max int=0
		left int=0
		right int=1
		size int =0
	)
	var freq [256]int

	for left<len(s){
		if right<len(s)&&freq[s[right]-'a']==0{
			freq[s[right]-'a']++
			right++
			size=right-left
			max=maxNum(max,size)
		}else{
			for freq[s[right]-'a']!=0{
				freq[s[left]-'a']--
				left++
				size=right-left
			}
			max=maxNum(max,size)
			//freq[s[right]-'a']++
			//right++
		}
	}
	return max

}
















