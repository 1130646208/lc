package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	var ret []int
	ret = append(ret, reversePrint(head.Next)...)
	ret = append(ret, head.Val)
	return ret
}
