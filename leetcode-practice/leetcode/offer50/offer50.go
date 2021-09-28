package offer50
//ac
func firstUniqChar(s string) byte {
	if len(s)==0{
		return ' '
	}
	var (
		ret byte=' '
	)
	m:= make(map[byte]int)
	for i:=0;i<len(s);i++{
		m[s[i]]++
	}
	for i:=0;i<len(s);i++{
		if m[s[i]]==1{
			ret=s[i]
			break
		}
	}
	return ret
}