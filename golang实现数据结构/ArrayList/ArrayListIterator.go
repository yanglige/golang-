package ArrayList

import "errors"

type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
	Remove()
	GetIndex() int

}

type Iterable interface {
	Iterator() Iterator // 初始化接口

}

// 构造指针访问数组
type ArrayListIterator struct {
	list *ArrayList
	currentindex int // 当前索引

}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator)  // 构造迭代器
	it.currentindex = 0
	it.list = list
	return it
}

func (it *ArrayListIterator) HasNext() bool{

	return it.currentindex < it.list.TheSize // 是否有下一个
}
func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("没有下一个")
	}
	value, err := it.list.Get(it.currentindex)
	it.currentindex++
	return value, err
}
func (it *ArrayListIterator) Remove() {
	it.currentindex--
	it.list.Delete(it.currentindex) // 删除一个元素

}
func (it *ArrayListIterator) GetIndex() int {
	return it.currentindex
}