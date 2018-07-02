package algorithm

/*
Input: s = "PAYPALISHIRING",
numRows = 3
Explanation1:
P   A   H   N
A P L S I I G
Y   I   R
Output:"PAHNAPLSIIGYIR"

numRows=4
Explanation2:
P     I    N
A   L S  I G
Y A   H R
P     I
Output: "PINALSIGYAHRPI"

解析：第一行和最后一行情况特殊，自己位于对角线两端，其后一个字符在原字符串中的位置满足loc := i+j*(2*rowNum-2)
每往下走一行，竖列坐标+1，对角线上坐标-1，合起来距离-2，故可以得出变化前后对角线上字母在原字符串中坐标 loc := i+j*(2*rowNum-2) - i*2
 */
func ZigZagConversion(str string, rowNum int) string {
	length := len(str)
	if rowNum == length || len(str) == 0 || rowNum > len(str) {
		return str
	}
	var ret []byte
	var loc int
	var baseGap = 2*rowNum - 2
	for i := 0; i < rowNum; i++ {
		j := 1
		ret = append(ret, str[i])
		for loc < length {
			if i == 0 || i == rowNum-1 {
				loc = i + j*(baseGap)
				if loc < length {
					ret = append(ret, str[loc])
				} else {
					break
				}
			} else {
				loc = i + j*(baseGap) - i*2
				if loc < length {
					ret = append(ret, str[loc])
				} else {
					break
				}
				loc = i + j*(baseGap)
				if (loc < length) {
					ret = append(ret, str[loc])
				} else {
					break
				}
			}
			j++
		}
	}
	return string(ret)
}
