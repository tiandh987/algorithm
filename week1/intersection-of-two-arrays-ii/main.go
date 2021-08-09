package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{4, 9, 5, 4, 4}
	nums2 := []int{9, 4, 9, 8, 4, 6}

	//res := intersect(nums1, nums2)
	res := intersect2(nums1, nums2)

	fmt.Println(res)
}

/**
 * leetcode 350
 * 两个数组的交集II
 * https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
 */

/**
 * 方法一： 哈希表法
 * 同一个数组在两个数组中可能会出现多次，需要用hash表存储每个数字出现的次数。
 * 对于一个数字，其在交集中出现的次数，等于该数字在两个数组中出现的最小次数。
 *
 * 首先遍历第一个数组，并在hash表中记录下数组中的数字，以及出现的次数；
 * 然后遍历第二个数组，如果数字在hash表中存在，则将该数字添加到“答案”，并将该数字的次数在hash表中减一。
 *
 * 为了降低空间复杂度，首先遍历较短的数组并在哈希表中记录每个数字以及对应出现的次数，然后遍历较长的数组得到交集。
 *
 * 复杂度分析
 * 时间复杂度：O(m+n)
 * 空间复杂度：O(min(m,n))
 */
func intersect(nums1 []int, nums2 []int) []int {
	// 该段代码用于：先遍历较短的数组。
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}

	intersection := []int{}
	for _, num := range nums2 {
		if m[num] > 0 {
			intersection = append(intersection, num)
			m[num]--
		}
	}
	return intersection
}

/**
 * 方法二： 排序 + 双指针
 * 1、排序。
 * 2、两个指针，分别指向两个数组首个元素。
 * 3、同时遍历，如果相等，则加入“答案”，并将两个指针同时向后移动一位
 * 4、如果不相等，则将数字小的指针向后移动一位
 * 5、直到一个数组先遍历结束
 *
 * 复杂度分析
 * 时间复杂度：O(mlogm + nlogn + m + n) 排序时间 + 遍历时间
 * 空间复杂度：O(min(m,n))
 */
func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	first := 0
	second := 0
	length1 := len(nums1)
	length2 := len(nums2)

	res := []int{}
	for first < length1 && second < length2 {
		if nums1[first] == nums2[second] {
			res = append(res, nums1[first])
			first++
			second++
		} else if nums1[first] < nums2[second] {
			first++
		} else if nums1[first] > nums2[second] {
			second++
		}
	}
	return res
}
