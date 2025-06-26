/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    nn := head.Next

    for nn != nil && nn.Next != nil {
        nn = nn.Next.Next
        head = head.Next
    }

    if nn != nil{
        return head.Next
    }
    return head
}