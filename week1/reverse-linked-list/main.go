package main

type ListNode strunct {
	Val int
	Next *ListNode
}

func main() {
	
}

// leetcode 206
// 翻转链表
// https://leetcode-cn.com/problems/reverse-linked-list/

// 思路：头插法
// 复杂度分析：
// 时间复杂度：O(n)
// 空间复杂度：O(1).
func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{}

	for head != nil {
		next := head.Next

		head.Next = dummy.Next
		dummy.Next = head
		head = next
	}

	return dummy.Next 
}

