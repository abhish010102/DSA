/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// var head *ListNode
	head := &ListNode{}
	nh := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			nh.Next = list1
			list1 = list1.Next
			nh = nh.Next
		} else {
			nh.Next = list2
			list2 = list2.Next
			nh = nh.Next
		}
	}

    for list1 != nil{
        nh.Next = list1
        list1 = list1.Next
        nh = nh.Next
    }

    for list2 != nil{
        nh.Next = list2
        list2 = list2.Next
        nh = nh.Next
    }

	return head.Next
}