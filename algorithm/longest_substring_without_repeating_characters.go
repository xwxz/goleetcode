package algorithm

/**
 * 给定一个字符串，找出不含有重复字符的最长子串的长度。
 * 示例：
 * 给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。
 * 给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。
 * 给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列  而不是子串。
 * 为了调试方便，这里顺带给出了最长子串
 * 此题解法是快慢指针，注意指针只能向前，不能回退
 */
func LongestSubstringWithoutRepeatingCharacters(str string) (int) {
	if len(str) < 2 {
		return len(str)
	}
	var characterMap = make(map[byte]int)
	var maxLength int
	low := 0
	high := 0
	for ; high < len(str); high++ {
		ch := str[high]
		idx, ok := characterMap[ch]
		if ok && low <= idx {
			low = idx + 1
		} else {
			currentLen := high - low + 1
			if maxLength < currentLen {
				maxLength = currentLen
			}
		}
		characterMap[ch] = high
	}
	return maxLength
}
