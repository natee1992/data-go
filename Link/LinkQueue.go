package Link

type QueueLink struct {
	rear  *Node
	front *Node
}

type LinkQueue interface {
	Enqueue(data interface{})
	Dequeue() interface{}
	Length() int
}

func NewLinkQueue() *QueueLink {
	return &QueueLink{}
}

func (qlk *QueueLink) Length() int {
	p := qlk.front
	length := 0
	for p.PNext != nil {
		p = p.PNext
		length++
	}
	return length

}

func (qlk *QueueLink) Enqueue(value interface{}) {
	newNode := &Node{value, nil}
	if qlk.front == nil {
		qlk.front = newNode //插入一个节点
		qlk.rear = newNode
	} else {
		qlk.rear.PNext = newNode
		qlk.rear = qlk.rear.PNext
	}
}

func (qlk *QueueLink) Dequeue() interface{} {
	if qlk.front == nil {
		return nil
	}
	newnode := qlk.front //记录头部位置
	if qlk.front == qlk.rear { //只剩下一个
		qlk.front = nil
		qlk.rear = nil
	} else {
		qlk.front = qlk.front.PNext
	}
	return newnode.Data
}
