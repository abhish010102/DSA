type ListNode1 struct {
	Val  int
	Next *ListNode1
}

type MyLinkedList struct {
	head *ListNode1
}

func Constructor() MyLinkedList {
	return MyLinkedList{head: nil}
}

func (this *MyLinkedList) Get(index int) int {
	h := this.head
    
    for index>0 {
        if h==nil {
            return -1
        }
        index--
        h=h.Next
    } 
    if(h!=nil){
        return h.Val
    }
    return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	nnode := &ListNode1{Val: val,Next:this.head}
	this.head = nnode
}

func (this *MyLinkedList) AddAtTail(val int) {
	nnode := &ListNode1{Val: val}
	if this.head == nil {
		this.head = nnode
        return
	} 
    h1:=this.head
    for h1.Next!=nil{
        h1=h1.Next
    }
    h1.Next=nnode;    
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
    nnode := &ListNode1{Val: val}
	if index==0 {
		nnode.Next = this.head
        this.head = nnode
        return
	} 
    if(this.head ==nil ){
        return
    }
    h1:=this.head
    for index!=1 && h1.Next!=nil {
        index--;
        h1=h1.Next
    }
    if(index == 1 && h1!=nil){
        nnode.Next= h1.Next
        h1.Next = nnode
    }
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
    if  this.head==nil{
        return
    }
    if index==0  {
        this.head = this.head.Next
        return
	} 
    h1:=this.head
    for index!=1 && h1.Next!=nil {
        index--;
        h1=h1.Next
    }
    if(index == 1) && h1.Next !=nil{
        h1.Next = h1.Next.Next
    }
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */