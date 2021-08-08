package main

func main() {

}

// leetcode 66
// 加一
// https://leetcode-cn.com/problems/plus-one/

// 思路：
// 45 + 1
// 125 + 5
// 999 + 1

// 复杂度分析
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func plusOne(digits []int) []int {
	for tail := len(digits) - 1; tail >= 0; tail-- {
		digits[tail]++
		digits[tail] %= 10
		if 0 != digits[tail] {
			return digits
		}
	}
	return append([]int{1}, digits...)
}
