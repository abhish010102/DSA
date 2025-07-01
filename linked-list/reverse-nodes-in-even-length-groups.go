/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseEvenLengthGroups(head *ListNode) *ListNode {
	h1 := head

	for i := 2; h1 != nil && h1.Next != nil; i++ {

		h2, index := h1.Next, 1
        
        for(h2!=nil && index<=i){
            h2=h2.Next
            index++;
        }

        
		if index%2 != 0 {
			h2 = h1.Next
			if h2.Next == nil {
				return head
			}

			l1 := h2.Next
			for j := 1; j < i && l1 != nil; j++ {
				l2 := l1.Next
				l1.Next = h2
				h2 = l1
				l1 = l2
			}
			h1.Next.Next = l1
			l1 = h1.Next
			h1.Next = h2
			h1 = l1
		} else {
			for j := 0; j < i && h1 != nil; j++ {
				h1 = h1.Next
			}
		}
	}

	return head
}