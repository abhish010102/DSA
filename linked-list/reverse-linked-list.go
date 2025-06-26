/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    // if head == nil || head.Next == nil{
    //     return head
    // }
    
    // l1 := reverseList(head.Next)
    // ret := l1
    // for (l1.Next != nil){
    //     l1=l1.Next
    // }

    // l1.Next=head
    // head.Next=nil
    // return ret

    if head == nil{
        return head
    }
    
    l1 := head.Next
    head.Next = nil

    for l1 != nil {
        l2 := l1.Next
        l1.Next = head
        head=l1
        l1 = l2
    } 

    return head
}