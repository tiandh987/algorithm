package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	//rotate1(nums, k)
	rotate2(nums, k)

	for _, value := range nums {
		fmt.Print(value)
	}
}

// leetcode 189
// 旋转数组
// https://leetcode-cn.com/problems/rotate-array/

// 思路：遍历一遍数组，将元素放在新数组正确的位置

// 复杂度分析
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func rotate1(nums []int, k int) {
	length := len(nums)

	if length == 0 {
		return
	}

	newArray := make([]int, length)

	for index, value := range nums {
		newIndex := (index + k) % length
		newArray[newIndex] = value
	}

	copy(nums, newArray)
}

func reverse(nums []int) {
	length := len(nums)
	for i := 0; i < (length / 2); i++ {
		nums[i], nums[length-i-1] = nums[length-i-1], nums[i]
	}
}

/**
 * 方法二：翻转数组
 * 操作	                             结果
 * 原始数组	                         1 2 3 4 5 6 7
 * 翻转所有元素                      7 6 5 4 3 2 1
 * 翻转[0, (k mod n) −1]区间的元素	 5 6 7 4 3 2 1
 * 翻转[(k mod n), (n−1)]区间的元素	 5 6 7 1 2 3 4
 */

func rotate2(nums []int, k int) {
	k %= len(nums)

	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
