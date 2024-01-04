/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 此题要画图,列个数学公式就行
// a=起点到交点距离
// b=交点到相遇点距离
// c=相遇点回到交点的距离
// 慢走的距离=a+b
// 快走的距离=a+b+c+b
// 慢走的距离*2=快走的距离
// 得a=c,即慢指针再从头一步一步走,快指针从相遇点一步一步走,就会在相交点相遇了
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}

	// 即慢指针再从头一步一步走
	slow = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}

// @lc code=end

