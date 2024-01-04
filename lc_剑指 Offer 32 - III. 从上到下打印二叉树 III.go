package main

func levelOrder3(root *TreeNode) [][]int {

	queue := []*TreeNode{root}
	var ret [][]int
	times := 0

	for len(queue) > 0 {
		times++
		l := len(queue)
		var level = make([]int, l)

		for i, n := range queue {
			// odd
			if times&1 == 1 {
				level[i] = n.Val
			} else {
				level[l-1-i] = n.Val
			}

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
