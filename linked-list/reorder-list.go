/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
    if(head==nil || head.Next==nil || head.Next.Next==nil){
        return
    }
    
	f := head
	s := head

	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
	}

	f = s.Next
	s.Next = nil
	s = f.Next
	f.Next = nil

	for s != nil {
		l1 := s.Next
		s.Next = f
		f = s
		s = l1
	}

	s = head

	for f != nil {
		l1 := s.Next
		s.Next = f
		f = f.Next
		s.Next.Next = l1
		s = l1
	}
}