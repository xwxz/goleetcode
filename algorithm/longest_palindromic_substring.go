package algorithm
//查找最长回文子串

//如 "abcbcbd" => "bcbcb"
func FindLongestPalindromicSubstring(str string) string {
	n := len(str)
	if n < 2 {
		return str
	}

	maxLen := 1
	startIdx := 0
	endIdx := 0
	for i := 1; i < n; {
		if n-i < maxLen/2 {//当剩余字符长度总和小于当前找到的最大字符串长度一半时，表示当前找到的就是最大的
			break
		}
		lo := i
		hi := i
		for hi < n-1 && str[hi] == str[hi+1] {//重复的字符可随意翻转，固可以认为其为一个字符
			hi++
		}
		i = hi + 1
		for lo > 0 && hi < n-1 && str[lo-1] == str[hi+1] {
			lo--
			hi++
		}
		newLen := hi - lo + 1
		if newLen > maxLen {
			startIdx = lo
			endIdx = hi
			maxLen = newLen
		}
	}
	return str[startIdx:endIdx+1]
}
