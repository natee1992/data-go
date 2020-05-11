package Queue

type MyQueue interface {
	Size() int                //大小
	Front() interface{}       //第一个元素
	End() interface{}         //最后一个元素
	IsEmpty() bool            // 是否为空
	EnQueue(data interface{}) //入队
	DeQueue() interface{}     //出队
	Clear()                   //清空
}

type Queue struct {
	dataStore []interface{} //队列的数据存储
	theSize   int           //队列的大小
}

func NewQueue() *Queue {
	myqueue := new(Queue) //初始化开辟结构体
	myqueue.dataStore = make([]interface{}, 0)
	myqueue.theSize = 0
	return myqueue
}

func (mq *Queue) Size() int {
	return mq.theSize
}
func (mq *Queue) Front() interface{} {
	if mq.theSize == 0 {
		return nil
	} else {
		return mq.dataStore[0]
	}
}

func (mq *Queue) End() interface{} {
	if mq.theSize == 0 {
		return nil
	} else {
		return mq.dataStore[mq.theSize-1]
	}
}

func (mq *Queue) IsEmpty() bool {
	if mq.theSize == 0 {
		return true
	} else {
		return false
	}
}

func (mq *Queue) EnQueue(data interface{}) {
	mq.dataStore = append(mq.dataStore, data)
	mq.theSize++
}

func (mq *Queue) DeQueue() interface{} {
	if mq.Size() == 0 {
		return nil
	}
	data := mq.dataStore[0]
	if mq.Size() > 1 {
		mq.dataStore = mq.dataStore[1:mq.Size()]
	} else {
		mq.dataStore = make([]interface{}, 0)
	}
	mq.theSize--
	return data
}
func (mq *Queue) Clear() {
	mq.dataStore = make([]interface{}, 0)
	mq.theSize = 0
}
