package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	res := threeSum(nums)

	fmt.Println(res)
}

// leetcode 15
// 三数之和
// https://leetcode-cn.com/problems/3sum/

// 思路：排序 + 双指针
// a + b + c = 0 转换为 a + b = -c
// 复杂度分析：
// 时间复杂度：O(n^2)，排序O(nlogn) + 遍历数组O(n) + 双指针O(n)。
//   总体为：O(nlogn) + O(n) * O(n) = O(n^2)
// 空间复杂度：O(1).
func threeSum(nums []int) [][]int {
	// 对切片中元素进行升序排序
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		// 因为已经对切片中数据排序，所以如果 nums[i] > 0,
		// 则a + b + c 必然大于0.
		if nums[i] > 0 {
			break
		}

		// 去重
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		head := i + 1
		tail := len(nums) - 1
		for head < tail {
			a := nums[head]
			b := nums[tail]
			c := nums[i]

			sum := a + b + c
			if sum == 0 {
				res = append(res, []int{a, b, c})
				// 去重
				for head < tail && a == nums[head] {
					head++
				}
				for head < tail && b == nums[tail] {
					tail--
				}
			} else if sum < 0 {
				head++
			} else {
				tail--
			}
		}
	}

	return res
}
