package codeLink

type QueueLink struct {
	//data interface{}
	rear *Node
	front *Node
}

type LinkQueue interface {
	length() int
	Enqueue(value interface{})
	Dequeue() interface{}

}
func NewLinkQueue() *QueueLink {
	return &QueueLink{}
}

func (qlk *QueueLink) length() int {
	pnext := qlk.front
	length := 0
	for pnext.pNext != nil {
		pnext = pnext.pNext
		length++
	}
	return length
}

func (qlk *QueueLink) Enqueue(value interface{}) {
	newnode := &Node{value, nil}
	if qlk.front == nil {
		qlk.front = newnode
		qlk.rear = newnode
	}else {
		qlk.rear.pNext = newnode
		qlk.rear = qlk.rear.pNext
	}
}
func (qlk *QueueLink) Dequeue() interface{} {
	if qlk.front == nil {
		return nil
	}
	newnode := qlk.front
	if qlk.front == qlk.rear {
		qlk.front = nil
		qlk.rear = nil
	}else {
		qlk.front = qlk.front.pNext
	}
	return newnode.data
}