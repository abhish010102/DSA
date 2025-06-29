/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func getDecimalValue(head *ListNode) int {
    var res int
    for head != nil {
        res = (res << 1) | head.Val
        head = head.Next
    }
    return res
}