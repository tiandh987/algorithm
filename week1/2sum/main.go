package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	res := towSum(nums, target)
	fmt.Println(res)
}

// leetcode 1
// 两数之和
// https://leetcode-cn.com/problems/two-sum/

// 思路：a + b = target  =>  b = target - a
// 复杂度分析：
// 时间复杂度：O(n)，数组遍历
// 空间复杂度：O(n)，存储每个元素
func towSum(nums []int, target int) []int {
	preNums := map[int]int{}
	res := []int{}

	for index, num := range nums {
		value := target - num
		valueIndex, ok := preNums[value]
		if !ok {
			preNums[num] = index
		} else {
			res = append(res, valueIndex, index)
		}
	}
	return res
}
