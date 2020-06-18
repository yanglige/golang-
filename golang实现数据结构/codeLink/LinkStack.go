package codeLink



type Node struct {
	data interface{}
	pNext *Node

}
type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}

func NewStack() *Node {
	return &Node{} // 返回一个节点指针
}

func (n *Node) IsEmpty() bool {
	if &n.pNext == nil {
		return true
	}else {
		return false
	}

}
func (n *Node) Push(data interface{}) {
	newnode := &Node{data, nil}
	newnode.pNext = n.pNext  // 头插
	n.pNext = newnode


}
func (n *Node) Pop() interface{} {
	if n.IsEmpty() == true {
		return nil
	}
	value := n.pNext.data
	n.pNext = n.pNext.pNext
	return value

}
func (n *Node) Length() int {
	pnext := n
	length := 0
	for pnext.pNext != nil {
		pnext = pnext.pNext
		length++
	}
	return length

}



