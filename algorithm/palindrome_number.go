package algorithm


/*
题目：
	判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例：
输入: 121
输出: true

示例：
输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

解析：
整数反转，我们可以利用除法求余的方法进行计算反转后的整数，如果反转后还一样，那说明就是回文整数，这里不用考虑溢出，溢出后自然就不会相等了
*/

func IsPalindromeNumber(num int) bool {
	if num < 0 {
		return false
	}
	var pNum int
	sourceNum := num
	for num > 0 {
		tail := num % 10
		pNum = pNum*10 + tail
		num = num / 10
	}
	return sourceNum == pNum
}
