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
	dataSource  [] interface{}
	capSize     int // 最大范围
	currentSize int //当前实际使用大小
}

func NewStack() *Stack {
	myStack := new(Stack)
	myStack.dataSource = make([]interface{}, 0, 10)
	myStack.capSize = 10
	myStack.currentSize = 0
	return myStack
}

func (myStack *Stack) Clear() {
	myStack.dataSource = make([]interface{}, 0, 10)
	myStack.capSize = 10
	myStack.currentSize = 0
}

func (myStack *Stack) Size() int {
	return myStack.currentSize
}
func (myStack *Stack) Pop() interface{} {

}
func (myStack *Stack) Push() {

}
func (myStack *Stack) IsEmpty() bool {
	if myStack.currentSize == 0 {
		return true
	} else {
		return false
	}
}
func (myStack *Stack) IsFull() bool {
	if myStack.currentSize >= myStack.capSize {
		return true
	} else {
		return false
	}
}
