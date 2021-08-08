package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	len := removeDuplicates(nums)
	fmt.Println("len: ", len)
	for i := 0; i < len; i++ {
		fmt.Print(nums[i])
	}
}

// leetcode 26
// 删除有序数组中的重复项
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

// 思路：
// 数组nums是有序的，所以对于任意 i<j，如果nums[i] == nums[j], 则对任意 i <= k <= j,
// 必有 nums[i] == nums[k] == nums[j],
// 即：相等的元素在数组中的下标一定是连续的。

// 复杂度分析
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeDuplicates(nums []int) int {

	// 如果数组长度为0 直接返回0
	if len(nums) == 0 {
		return 0
	}

	// 数组长度大于0，则至少包含一个元素，在删除重复元素之后也至少剩下一个元素
	// 所以nums[0]保持原装即可，从下标1开始删除重复元素。
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
