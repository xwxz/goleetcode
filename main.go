package main

import (
	"github.com/xwxz/goleetcodes/structure"
	"github.com/xwxz/goleetcodes/algorithm"
	"fmt"
)

func testTwoSum() {
	arr := []int{1, 2, 3, 5}
	x, y := algorithm.TwoSum(arr, 4)
	fmt.Printf("%d,%d\n", x, y)
}

func testAddTwoNumbers() {
	list1 := new(structure.List)
	list1.Append(&structure.Node{Data: 2})
	list1.Append(&structure.Node{Data: 4})
	list1.Append(&structure.Node{Data: 3})
	list1.Print()

	list2 := new(structure.List)
	list2.Append(&structure.Node{Data: 5})
	list2.Append(&structure.Node{Data: 6})
	list2.Append(&structure.Node{Data: 4})
	list2.Print()

	list3 := algorithm.AddTwoNumbers(list1, list2)
	list3.Print()
}

func testLSWRC() {
	l, substring := algorithm.LongestSubstringWithoutRepeatingCharacters("Notethattheanswer")
	fmt.Printf("%s:%d\n", substring, l)
}

func testFindMedian() {
	var a = []int{1, 4, 6}
	var b = []int{2, 5}
	c := algorithm.MedianOfTwoSortedArrays(a, b)
	fmt.Println(c)
}

func testFindLongetstPalindrom() {
	fmt.Println(algorithm.FindLongestPalindromicSubstring("abcbcbd"))
}

func main() {
	testTwoSum()
	testAddTwoNumbers()
	testLSWRC()
	testFindMedian()
	testFindLongetstPalindrom()
}
