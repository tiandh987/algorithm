package main

func main() {
}

// leetcode 88
// 合并两个有序数组
// https://leetcode-cn.com/problems/merge-sorted-array/

// 复杂度分析
// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	tail := m + n - 1
	first := m - 1
	second := n - 1

	for first >= 0 || second >= 0 {
		if first < 0 {
			nums1[tail] = nums2[second]
			second--
		} else if second < 0 {
			nums1[tail] = nums1[first]
			first--
		} else {
			if nums1[first] >= nums2[second] {
				nums1[tail] = nums1[first]
				first--
			} else {
				nums1[tail] = nums2[second]
				second--
			}
		}
		tail--
	}
	return
}
