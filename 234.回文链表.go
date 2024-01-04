/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

var newLasthalfHead *ListNode

func isPalindrome(head *ListNode) bool {

	// 如何用o(1)的空间和o(n)的时间?
	// 快慢指针找到中间节点,再把中间节点之后的链表反转,再判断
	var fast, slow = head, head
	for ; fast.Next != nil && fast.Next.Next != nil; fast = fast.Next.Next {
		slow = slow.Next
	}
	lastHalf := reverse(slow.Next)
	slow.Next = lastHalf

	for ; lastHalf != nil; lastHalf = lastHalf.Next {
		if head.Val != lastHalf.Val {
			return false
		}
		head = head.Next
	}

	return true
}

// 翻转链表
func reverse(node *ListNode) *ListNode {
	var preNode *ListNode
	var nextNode *ListNode
	for cur := node; cur != nil; cur = nextNode {
		nextNode = cur.Next
		cur.Next = preNode
		preNode = cur
	}
	return preNode
}

// @lc code=end

