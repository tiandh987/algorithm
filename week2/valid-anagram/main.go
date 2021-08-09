package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"

	res := isAnagram(s, t)

	fmt.Println(res)
}

/**
 * leetcode 242
 * 有效的字母异位词
 * https://leetcode-cn.com/problems/valid-anagram/
 *
 * 若 s 和 t 中每个字符出现的次数都相同，则称 s 与 t 互为字母异位词。
 *
 * 思路：哈希法
 *
 * 复杂度分析
 * 时间复杂度：O(n)
 * 空间复杂度：O(s), s 为字符集的长度，这儿为26
 */
func isAnagram(s string, t string) bool {
	var a1, a2 [26]int

	for _, ch := range s {
		a1[ch-'a']++
	}

	for _, ch := range t {
		a2[ch-'a']++
	}

	return a1 == a2
}

/**
 * 进阶：字符串中包含unicode字符
 */
func isAnagram2(s string, t string) {
	if len(s) != len(t) {
		return false
	}

	m := map[rune]int{}
	for _, ch := range s {
		m[ch]++
	}

	for _, ch := range t {
		m[ch]--
		if m[ch] < 0 {
			return false
		}
	}
	return true
}
