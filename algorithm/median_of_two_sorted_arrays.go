package algorithm

import "math"

//采用分治策略算法，其中为了解决虚拟数组长度奇偶不定性，做虚拟插值让数组长度恒为奇数，给定一个整数n，则2n+1一定为奇数

//两个排序数组A和B，求中位数k上的值，则可以假设从A中取c1个数，从B中取c2个数，c1+c2=k且在c1和c2左边的数均小于其右边的数

//将A用c1分为两部分，左边叫做l1，右边叫做r1，将B用c2分为两部分，左边叫做l2，右边叫做r2，需要满足

//len(A[0...l1]) + len(B[0...l2]) == k 且 A[l1] < B[r2] && A[r1] > B[l1]

//则中位数median = (max(l1,l2) + min(r1,r2)) / 2

//其中为了提升效率，使用数组元素少的数组作为基准数组用于二分

//其中当c1为0或者c2为0时，表示某个数组所有值均处于中位数的同一边（c1为0，A在中位数右边，c2为0，B在中位数右边）
func MedianOfTwoSortedArrays(A, B []int) float64 {
	m := len(A)
	n := len(B)
	if m > n {
		return MedianOfTwoSortedArrays(B, A)
	}
	var l1, l2, r1, r2, c1, c2 int
	var l0 = 0
	var hi = 2 * m
	for l0 < hi {

		c1 = (hi - l0) / 2
		c2 = m + n - c1

		if c1 == 0 {
			l1 = math.MinInt64
		} else {
			l1 = A[(c1-1)/2]
		}

		if c1 == m*2 {
			r1 = math.MaxInt64
		} else {
			r1 = A[c1/2]
		}

		if c2 == 0 {
			l2 = math.MinInt64
		} else {
			l2 = B[(c2-1)/2]
		}

		if c2 == n*2 {
			r2 = math.MaxInt64
		} else {
			r2 = B[c2/2]
		}

		if l1 > r2 {
			hi = c1 - 1
		} else if l2 > r1 {
			l0 = c1 + 1
		} else {
			break
		}
	}

	return (math.Max(float64(l1), float64(l2)) + math.Min(float64(r1), float64(r2))) / 2
}
