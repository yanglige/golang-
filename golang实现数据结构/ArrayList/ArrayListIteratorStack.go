package ArrayList
type StackArrayX interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool

}


type StackX struct {
	myarray *ArrayList
	Myit Iterator
}

func NewArrayListStackX() *StackX {
	mystack := new(StackX)
	mystack.myarray = NewArrayList()
	mystack.Myit = mystack.myarray.Iterator()
	return mystack
}


func (mystack *StackX) Clear() {
	mystack.myarray.Clear()
	mystack.myarray.TheSize = 0


}
func (mystack *StackX) Size() int {
	return mystack.myarray.TheSize

}
func (mystack *StackX) Pop() interface{} {
	if !mystack.IsEmpty() {
		last := mystack.myarray.dataStore[mystack.myarray.TheSize-1]
		mystack.myarray.Delete(mystack.myarray.TheSize-1)
		return last
	}
	return nil

}
func (mystack *StackX) Push(data interface{}) {
	if !mystack.IsFull() {
		mystack.myarray.Append(data)
	}
}
func (mystack *StackX) IsFull() bool {
	if mystack.myarray.TheSize >= 10 {
		return true
	}else {
		return false
	}

}

func (mystack *StackX) IsEmpty() bool {
	if mystack.myarray.TheSize == 0 {
		return true
	}else {
		return false
	}

}

