/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

var newHead *ListNode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	reverse(head)
	return newHead
}

func reverse(node *ListNode) *ListNode {
	if node.Next == nil {
		newHead = node
		return node
	}
	nextNode := reverse(node.Next)
	nextNode.Next = node
	node.Next = nil
	return node
}

// @lc code=end

