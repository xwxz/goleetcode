package algorithm

/*
给定一个 32 位有符号整数，将整数中的数字进行反转。
示例1：
输入: 123
输出: 321
示例2：
输入: -123
输出: -321
示例3：
输入: 120
输出: 21
对于反转后超过数值范围，直接返回0

解析：
最开始有两个想法，一、转字符串进行按字符反转，二、取模进位加法
方法一有点麻烦，这里就不说了，下面给出的是方法二。
一个数除以一个数，不会溢出，结果会向0靠拢，不管除数是正是负，但是乘法会产生溢出，对于如果检测一个数乘以另一个数得到的结果是否溢出，
很简单，只需要反过来再除下，如果乘积除以其中一个乘数等于另一个乘数，则表示没有溢出。

*/
func ReverseInteger(x int32) int32 {
	var y int32
	for x != 0 {
		remainder := x % 10
		newRet := y*10 + remainder
		if (newRet-remainder)/10 != y {//判断是否溢出
			return 0
		}
		y = newRet
		x = x / 10
	}
	return int32(y)
}
