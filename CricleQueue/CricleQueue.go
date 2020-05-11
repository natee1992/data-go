package CricleQueue

import "errors"

const QueueSize = 5 //最多存储QueueSize -1
// 空一个空格表示满格

type CircleQueue struct {
	data  [QueueSize]interface{} //存储数据结构
	front int                    // 头部位置
	rear  int                    // 尾部位置
}

// initQueue: 头尾重合，为空
func InitQueue(q *CircleQueue) {
	q.front = 0
	q.rear = 0
}

// EnQueue: 入队
func EnQueue(q *CircleQueue, data interface{}) (err error) {
	if (q.rear+1)%QueueSize == q.front%QueueSize {
		return errors.New("队列已经满了")
	}
	q.data[q.rear] = data
	q.rear = (q.rear + 1) % QueueSize
	return nil
}

// DeQueue： 出队
func DeQueue(q *CircleQueue) (data interface{}, err error) {
	if q.rear == q.front {
		return nil, errors.New("队列为空")
	}
	res := q.data[q.front] //取出第一个数据
	q.data[q.front] = 0
	q.front = (q.front + 1) % QueueSize
	return res, nil
}

// QueueLength: 长度
func QueueLength(q *CircleQueue) int {
	return (q.rear - q.front + QueueSize) % QueueSize
}
