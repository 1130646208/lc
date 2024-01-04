package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{}

	if root != nil {
		queue = append(queue, root)
	}

	var ret [][]int

	for len(queue) > 0 {
		var level []int
		l := len(queue)

		for _, n := range queue {
			level = append(level, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		queue = queue[l:]
		ret = append(ret, level)
	}
	return ret
}
