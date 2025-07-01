type Node struct {
	Url  string
	Pre  *Node
	Next *Node
}

type BrowserHistory struct {
	head *Node
	cur  *Node
}

func Constructor(homepage string) BrowserHistory {
    head:=&Node{Url: homepage}
    head.Pre = head
	return BrowserHistory{head:head , cur: head}
}

func (this *BrowserHistory) Visit(url string) {
	h1 := &Node{Url: url, Pre: this.cur}
	this.cur.Next = h1
    this.head.Pre = h1
	this.cur = this.cur.Next
}

func (this *BrowserHistory) Back(steps int) string {
    for steps > 0 && this.head != this.cur {
        this.cur = this.cur.Pre
		steps--
	}
    return this.cur.Url
}

func (this *BrowserHistory) Forward(steps int) string {
	for steps > 0 && this.cur.Next != nil {
    	this.cur = this.cur.Next
		steps--
	}
	return this.cur.Url
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */



//  ----------------------------------------------------------------------------------
/*
type BrowserHistory struct {
    home string
    hist []string
    cur int
}

func Constructor(homepage string) BrowserHistory {
    bb := BrowserHistory{home: homepage}
    bb.hist = append(bb.hist, homepage)
    return bb
}

func (this *BrowserHistory) Visit(url string) {
	this.hist = this.hist[:this.cur+1]
    this.hist = append(this.hist, url)
    this.cur = len(this.hist)-1
}

func (this *BrowserHistory) Back(steps int) string {
    if steps < this.cur {
        this.cur -= steps
    } else {
        this.cur = 0
    }
    return this.hist[this.cur]
}

func (this *BrowserHistory) Forward(steps int) string {
	if steps < (len(this.hist)-1) - this.cur {
        this.cur += steps
    } else {
        this.cur = len(this.hist)-1
    }
    return this.hist[this.cur]
}

  
 */