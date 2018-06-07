package main

import (
	"github.com/xwxz/goleetcodes/twosum"
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 5}
	x, y := twosum.TwoSum(arr, 4)
	fmt.Printf("%d,%d", x, y)
}
