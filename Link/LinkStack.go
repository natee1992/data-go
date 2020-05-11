package Link

type Node struct {
	Data  interface{}
	PNext *Node
}

type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}

func NewStack() *Node {
	return &Node{} //返回一个节点指针
}

func (n *Node) IsEmpty() bool {
	if n.PNext == nil {
		return true
	} else {
		return false
	}
}

func (n *Node) Push(data interface{}) {
	newNode := &Node{}
	newNode.Data = data
	newNode.PNext = n.PNext
	n.PNext = newNode //头部插入

}

func (n *Node) Pop() interface{} {
	if n.IsEmpty() == true {
		return nil
	}
	value := n.PNext.Data //要弹出的数据
	n.PNext = n.PNext.PNext
	return value
}

func (n *Node) Length() int {
	p := n
	length := 0
	for p.PNext != nil {
		p = p.PNext
		length++
	}
	return length

}
