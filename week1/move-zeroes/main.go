package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println("Before: ", nums)
	moveZeroes(nums)
	fmt.Println(" After: ", nums)

}

// leetcode 283
// 移动零
// https://leetcode-cn.com/problems/move-zeroes/

// 思路：将整个数组分为两个区间，非零区域、零区域

// 复杂度分析：
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func moveZeroes(nums []int) {
	i := 0

	for j := 0; j < len(nums); j++ {
		// 非零值移动到非零区域
		if nums[j] != 0 {
			if nums[i] != nums[j] {
				nums[i] = nums[j]
				nums[j] = 0
			}
			i++
		}
	}
}
