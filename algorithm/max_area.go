package algorithm

// 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
// 画 n 条垂直线，使得垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 注意：你不能倾斜容器，n 至少是2

// 解析:采用高低指针
// 两条垂直线与横轴组成四边形，能容纳水的多少取决于最短的垂直线，能容纳的水转换为面积来计算，面积等于底边乘以短高假设两条线为Xi和Xj
// 设低指针i,高指正j, area = MIN(Xi,Xj) * (j-i)，当j-i减小时，为了让面积尽量少的减少，那么xi和xj应该往大的方向增加，如果X[i]<X[i-1],那么直接跳过，
// 因为底边减小，高减小，面积肯定减小
//
func MaxArea(height []int) int {
	var low, hi int
	low = 0
	hi = len(height) - 1
	maxArea := 0
	for low < hi {
		if height[hi] > height[low] {
			if low-1 > 0 && height[low] < height[low-1] {
				low++
				continue
			}
			area := height[low] * (hi - low)
			if maxArea < area {
				maxArea = area
			}
			low++
		} else {
			if hi < len(height)-1 && height[hi] < height[hi+1] {
				hi--
				continue
			}
			area := height[hi] * (hi - low)
			if maxArea < area {
				maxArea = area
			}
			hi--
		}
	}
	return maxArea
}
