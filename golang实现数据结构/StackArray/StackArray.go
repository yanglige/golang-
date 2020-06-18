package StackArray

type StackArray interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool

}


type Stack struct {
	dataSource [] interface{}
	capsize int
	currentsize int

}

func NewStack() *Stack {
	mystack := new(Stack)
	mystack.dataSource = make([]interface{}, 0, 10)
	mystack.capsize = 10
	mystack.currentsize = 0
	return mystack
}


func (mystack *Stack) Clear() {
	mystack.dataSource = make([]interface{}, 0, 10)
	mystack.currentsize = 0

}
func (mystack *Stack) Size() int {
	return mystack.currentsize

}
func (mystack *Stack) Pop() interface{} {
	if !mystack.IsEmpty() {
		last := mystack.dataSource[mystack.currentsize-1]

		mystack.dataSource = mystack.dataSource[:mystack.currentsize-1]
		mystack.currentsize--
		return last
	}
	return nil

}
func (mystack *Stack) Push(data interface{}) {
	if !mystack.IsFull() {
		mystack.dataSource = append(mystack.dataSource, data)
		mystack.currentsize++
	}
}
func (mystack *Stack) IsFull() bool {
	if mystack.currentsize >= mystack.capsize {
		return true
	}else {
		return false
	}

}

func (mystack *Stack) IsEmpty() bool {
	if mystack.currentsize == 0 {
		return true
	}else {
		return false
	}
	
}


