/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	l1 := map[*Node]*Node{}
	nh := &Node{}
	h1, cur := head, nh

	for h1 != nil {
		cur.Next = &Node{Val: h1.Val}
		l1[h1] = cur.Next
		h1 = h1.Next
		cur = cur.Next
	}

	h1 = head
	for h1 != nil {
		l1[h1].Random = l1[h1.Random]
		h1 = h1.Next
	}
	return nh.Next
}