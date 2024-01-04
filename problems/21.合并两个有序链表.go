/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var dumbHead = &ListNode{Next: list1}

	pre, cur := dumbHead, list1

	// list2 != nil 很关键
	for cur != nil && list2 != nil {
		if cur.Val >= list2.Val {
			tmp := list2.Next
			list2.Next = cur
			pre.Next = list2
			list2 = tmp
			pre = pre.Next
		} else {
			cur = cur.Next
			pre = pre.Next
		}
	}

	if list2 != nil {
		pre.Next = list2
	}
	return dumbHead.Next
}

// @lc code=end

