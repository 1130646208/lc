package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	var newHead = head

	for head != nil {
		m[head] = &Node{
			Val:    head.Val,
			Next:   nil,
			Random: nil,
		}
		head = head.Next
	}

	for k, v := range m {
		v.Next = m[k.Next]
		v.Random = m[k.Random]
	}

	return newHead
}

func main() {
	copyRandomList(nil)
}
