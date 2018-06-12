package algorithm

import "github.com/xwxz/goleetcodes/structure"

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
