/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// using two pointers
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	h1, h2 := headA, headB
// 	for h1 != nil {
// 		h2 = headB
//         for h2 != nil  {
//             if h1==h2{
//                 return h1
//             }
//             h2 = h2.Next
//         }
//         h1= h1.Next
// 	}
// 	return nil
// }

// use map
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ml := make(map[*ListNode]bool)
	for h1:=headA; h1 != nil;h1=h1.Next {
        ml[h1] = true
	}
	for h1:=headB; h1 != nil;h1=h1.Next {
        if ml[h1] == true{
            return h1
        }
	}
	return nil
}
