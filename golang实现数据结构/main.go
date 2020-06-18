package main

import (
	"fmt"
	"golang实现数据结构/ArrayList"
	"golang实现数据结构/CircleQueue"
	"golang实现数据结构/Queue"
	"golang实现数据结构/StackArray"
	"golang实现数据结构/codeLink"
)

func main1()  {
	list := ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	fmt.Println(list)
}

func main2()  {
	list := ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("a2")
	list.Append("a3")
	fmt.Println(list.TheSize)
}
func main3()  {
	list := ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("a2")
	list.Append("a3")
	fmt.Println(list.TheSize)
}
func main4()  {
	list := ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("a2")
	list.Append("a3")
	for i := 0; i < 20; i++ {
		list.Insert(1,"1")
		fmt.Println(list)
	}
	list.Delete(0)
	fmt.Println(list)
	fmt.Println(list.TheSize)
}
func main5()  {
	list := ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("a2")
	list.Append("a3")
	for it := list.Iterator(); it.HasNext(); {
		item,_ := it.Next()
		if item == "a2" {
			it.Remove()
		}
	}
	fmt.Println(list)
	fmt.Println(list.TheSize)
}
func main6()  {
	mystack := StackArray.NewStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	mystack.Push(4)
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
}
func main7()  {
	mystack := ArrayList.NewArrayListStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	mystack.Push(4)
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
}
func main8()  {
	mystack := ArrayList.NewArrayListStackX()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	mystack.Push(4)

	for it := mystack.Myit; it.HasNext(); {
		item,_ := it.Next()
		if item == 2 {
			it.Remove()
		}
	}
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
}

func main9() {
	myq := Queue.NewQueue()
	myq.EnQueue(1)
	myq.EnQueue(2)
	myq.EnQueue(3)
	myq.EnQueue(4)
	myq.EnQueue(5)
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	myq.EnQueue(5)
	fmt.Println(myq.DeQueue())
}
func main10() {
	var myq CircleQueue.CricleQueue
	CircleQueue.InitQueue(&myq)
	CircleQueue.EnQueue(&myq, 1)
	CircleQueue.EnQueue(&myq, 2)
	CircleQueue.EnQueue(&myq, 3)
	CircleQueue.EnQueue(&myq, 4)
	CircleQueue.EnQueue(&myq, 5)
	CircleQueue.EnQueue(&myq, 6)
	fmt.Println(CircleQueue.DeQueue(&myq))
	fmt.Println(CircleQueue.DeQueue(&myq))
	fmt.Println(CircleQueue.DeQueue(&myq))
	fmt.Println(CircleQueue.DeQueue(&myq))
	fmt.Println(CircleQueue.DeQueue(&myq))


}

// 数组存储 修改查找O(1)  删除插入 O(n) 内存中连续存储   计算机理论上不存在大片连续内存

// 链表 删除和插入O(1)  查找修改O(n)
func main11() {
	mystack := codeLink.NewStack()
	for i:=0; i<10000000;i++ {
		mystack.Push(i)
	}
	for data := mystack.Pop();data != nil; data=mystack.Pop() {
		fmt.Println(data)
	}
}
func main() {
	mystack := codeLink.NewLinkQueue()
	for i:=0; i<10000000; i++ {
		mystack.Enqueue(i)
	}
	for data := mystack.Dequeue();data != nil; data=mystack.Dequeue() {
		fmt.Println(data)
	}
}