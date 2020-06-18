package CircleQueue

import "errors"

const QueueSize = 100 // 只能存放queuesize-1个数据
type CricleQueue struct {
	 data [QueueSize] interface{}
	 front int // 头部位置
	 rear int  // 尾部位置
}


func InitQueue(q *CricleQueue) {
	q.front = 0
	q.rear = 0
}


func Queuelength(q *CricleQueue) int {
	return (q.rear - q.front + QueueSize) % QueueSize
}
func EnQueue(q *CricleQueue, data interface{}) (err error) {
	if (q.rear+1) % QueueSize == q.front % QueueSize {
		return errors.New("队列已经满了")

	}
	q.data[q.rear] = data
	q.rear = (q.rear+1) % QueueSize
	return nil

}
func DeQueue(q *CricleQueue) (data interface{}, err error) {
	if q.rear == q.front {
		return nil, errors.New("队列为空")
	}
	res := q.data[q.front] // 取出第一个数据
	q.data[q.front] = 0 // 清空数据
	q.front = (q.front + 1) % QueueSize
	return res, err
}








