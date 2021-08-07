package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	area := maxArea(height)

	fmt.Println("maxArea : ", area)
}

// leetcode 11
// 盛水做多的容器
// https://leetcode-cn.com/problems/container-with-most-water/

// 思路：双指针
// area = width * minHeight
// 容器的盛水量取决于左右垂线中较小的
// 随着左右指针不断向内收，width会逐渐变小，所以minHeight应该尽可能的大

func maxArea(height []int) int {
	maxArea := 0
	head := 0
	tail := len(height) - 1

	for head < tail {
		width := tail - head
		minHeight := 0
		if height[head] < height[tail] {
			minHeight = height[head]
			head++
		} else {
			minHeight = height[tail]
			tail--
		}

		area := width * minHeight
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
