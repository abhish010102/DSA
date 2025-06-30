/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	var inner func(root *Node) *Node

	inner = func(root *Node) *Node {
		h1, h2 := root, root

		for h1 != nil {
			if h1.Child == nil {
				h2 = h1
				h1 = h1.Next
			} else {
				c := inner(h1.Child)
				c.Next = h1.Next
				if h1.Next == nil {
					h1.Next = h1.Child
				} else {
					h1.Next.Prev = c
					h1.Next = h1.Child
				}
				h1.Child.Prev = h1
				h1.Child = nil
			}
		}
		return h2
	}
	inner(root)
	return root
}