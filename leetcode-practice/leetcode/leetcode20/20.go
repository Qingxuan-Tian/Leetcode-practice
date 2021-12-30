package leetcode20

//栈做，ac
func isValid(s string) bool {
	hash := map[byte]byte{'(':')', '{':'}','[':']'}
	var stack []byte
	for i:=0;i<len(s);i++{
		if s[i]=='('||s[i]=='{'||s[i]=='['{
			stack=append(stack, hash[s[i]])
		}else {
			if len(stack)==0{
				return false
			}
			if stack[len(stack)-1]==s[i]{
				stack=stack[:len(stack)-1]
			}else {
				return false
			}
		}
	}
	//最后再检查一下栈中是否还有剩余的没有右括号可以配对的左括号
	if len(stack)==0{
		return true
	}else{
		return false
	}
}
