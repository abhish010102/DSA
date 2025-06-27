/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    for head!=nil && head.Val == val {
        head= head.Next
    }
    
    if head == nil {
        return head
    }

    h1:=head

    for h1.Next!=nil {
        if h1.Next.Val==val {
            h1.Next=h1.Next.Next;
        }else{
            h1=h1.Next;
        }
    }
    
    return head
}