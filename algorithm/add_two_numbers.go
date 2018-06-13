package algorithm

import "github.com/xwxz/goleetcodes/structure"

/**
 * 给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。
 * 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
 *
 * 示例：
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
 *
 */
func AddTwoNumbers(list1, list2 *structure.List) *structure.List {
	var list3 *structure.List
	if list1.Size >= list2.Size {
		list3 = list1
	} else {
		list3 = list2
	}

	p1 := list1.Head
	p2 := list2.Head
	p3 := list3.Head

	var sum, carry int
	for p1 != nil && p2 != nil {
		sum = p1.Data + p2.Data + carry
		carry = sum / 10
		p3.Data = sum % 10
		p1 = p1.Next
		p2 = p2.Next
		p3 = p3.Next
	}
	if carry != 0 {
		for carry > 0 && p3 != nil {
			sum = p3.Data + carry
			carry = sum / 10
			p3.Data = sum % 10
			p3 = p3.Next
		}
		if carry != 0 {
			list3.Append(&structure.Node{Data: carry})
		}
	}
	return list3
}
