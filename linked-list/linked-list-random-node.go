/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	head *ListNode
	l    int
}

func Constructor(head *ListNode) Solution {
	h1, c := head, 0
	for h1 != nil {
		c++
		h1 = h1.Next
	}
	return Solution{head: head, l: c}
}

func (this *Solution) GetRandom() int {
	h1, rn := this.head, rand.Intn(this.l)
	for rn > 0 {
		rn--
		h1 = h1.Next
	}
	return h1.Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */