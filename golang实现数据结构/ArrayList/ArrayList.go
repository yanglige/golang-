package ArrayList

import (
	"errors"
	"fmt"
)

// 接口
type List interface {
	Size() int
	Get(index int) (interface{}, error)
	Set(index int, newval interface{}) error
	Insert(index int, newval interface{}) error
	Append(newval interface{})
	Clear()
	Delete(index int) error
	string() string
}

type ArrayList struct {
	dataStore [] interface{}
	TheSize int
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.dataStore = make([]interface{}, 0, 10)
	list.TheSize = 0
	return list
}


func (list *ArrayList) Size() int {
	return list.TheSize
}

func (list *ArrayList) Get(index int) (interface{}, error){
	if index < 0 || index >= list.TheSize{
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Append(newval interface{}){
	list.dataStore = append(list.dataStore, newval)
	list.TheSize++
}


func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}


func (list *ArrayList) checkIsFull() {
	if list.TheSize == cap(list.dataStore) {
		newdataStore := make([]interface{}, 2*list.TheSize, 2*list.TheSize)
		copy(newdataStore, list.dataStore)
		list.dataStore = newdataStore
	}
}
func (list *ArrayList) Insert(index int, newval interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.checkIsFull()
	list.dataStore = list.dataStore[:list.TheSize+1]
	for i := list.TheSize; i > index; i-- {
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newval
	list.TheSize++

	return nil

}
func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)
	list.TheSize = 0
}
func (list *ArrayList) Set(index int, newval interface{}) error {
    if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
    list.dataStore[index] = newval
    return nil
}

func (list *ArrayList) Delete(index int) error {
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)
	list.TheSize--
	return nil
}