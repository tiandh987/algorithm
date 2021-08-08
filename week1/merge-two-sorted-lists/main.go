package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode 21
// 合并两个有序链表
// https://leetcode-cn.com/problems/merge-two-sorted-lists/

// 复杂度分析
// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 == nil {
		cur.Next = l2
	}

	if l2 == nil {
		cur.Next = l1
	}
	return dummy.Next
}
