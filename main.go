package main

import (
	"github.com/xwxz/goleetcodes/structure"
	"github.com/xwxz/goleetcodes/algorithm"
	"fmt"
	"math"
)

func testTwoSum() {
	arr := []int{1, 2, 3, 5}
	x := algorithm.TwoSum(arr, 4)
	fmt.Printf("%v\n", x)
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
	l := algorithm.LongestSubstringWithoutRepeatingCharacters("aab")
	fmt.Printf("%d\n", l)
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

func testFindZigZagConversion() {
	fmt.Printf("%s", algorithm.ZigZagConversion("PAYPALISHIRING", 4))
}

func testReverseInteger() {
	fmt.Println(algorithm.ReverseInteger(math.MaxInt32))
	fmt.Println(algorithm.ReverseInteger(math.MinInt32))
	fmt.Println(algorithm.ReverseInteger(-234))
	fmt.Println(algorithm.ReverseInteger(230))
}

func testAtoi() {
	fmt.Println(algorithm.Atoi("   -42"))
	fmt.Println(algorithm.Atoi("4193 with words"))
	fmt.Println(algorithm.Atoi("words and 987"))
	fmt.Println(algorithm.Atoi("-91283472332"))
	fmt.Println(algorithm.Atoi("3787791283472332"))
	fmt.Println(algorithm.Atoi("   "))
	fmt.Println(algorithm.Atoi("34 45 6l34   "))
	fmt.Println(algorithm.Atoi("2147483648"))
	fmt.Println(algorithm.Atoi("-91283472332"))
	fmt.Println(algorithm.Atoi("4193 with words"))
	fmt.Println(algorithm.Atoi("  0000000000012345678"))
	fmt.Println(algorithm.Atoi(" 0 123"))
	fmt.Println(algorithm.Atoi(" +0 123"))
	fmt.Println(algorithm.Atoi("-000000000000001"))
	fmt.Println(algorithm.Atoi("2147483648"))
}

func testPN() {
	fmt.Println(algorithm.IsPalindromeNumber(-121))
	fmt.Println(algorithm.IsPalindromeNumber(10))
	fmt.Println(algorithm.IsPalindromeNumber(141))
}

func testMaxArea() {

	area := algorithm.MaxArea([]int{1, 4, 7, 9, 2})
	fmt.Printf("maxArea = %d", area)
}

func testI2R() {
	ret := algorithm.IntegerToRoman(1994)
	fmt.Println(ret)
}

func main() {
	testLSWRC()
	testTwoSum()
	testAddTwoNumbers()
	testFindMedian()
	testFindLongetstPalindrom()
	testFindZigZagConversion()
	testReverseInteger()
	testAtoi()
	testPN()
	testMaxArea()
	testI2R()
}
