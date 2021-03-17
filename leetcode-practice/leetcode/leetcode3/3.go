package leetcode3

import "strings"

func lengthOfLongestSubstring(s string) int {
	var l,r,res=0,-1,0
	var freq[256] int
	var len = strings.Count(s,"")-1
	for {
		if r+1<len && freq[s[r+1]]==0{
			r++
			freq[s[r]]++
		}else {
			if(len==0){//考虑特殊情况，数组长度为0
				res=0
			}else{
				freq[s[l]]--
				l++
			}
		}
		if res<r-l+1{
			res=r-l+1
		}
		if l>=len{
			break
		}
	}
	return res
}
