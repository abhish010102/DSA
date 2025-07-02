/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    h1 := &ListNode{}
	c1,c2 := h1, h1

	for head != nil && head.Next !=nil {
        c2=head
        c1.Next=head.Next
        head=head.Next.Next
        c1.Next.Next=c2
        c1=c2
	}
    if(head!=nil){
	    c1.Next=head
    }else{
	    c1.Next=nil
    }
	return h1.Next
}