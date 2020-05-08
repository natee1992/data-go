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
	myStack.dataSource = make([]interface{}, 0, 5000)
	myStack.capSize = 5000
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
	if !myStack.IsEmpty(){
		last:=myStack.dataSource[myStack.currentSize-1]// 最后一个数据
		myStack.dataSource=myStack.dataSource[:myStack.currentSize-1] //删除最后一个
		myStack.currentSize-- //删除
		return last
	}
	return nil
}
func (myStack *Stack) Push(data interface{}) {
	if !myStack.IsFull(){
		myStack.dataSource = append(myStack.dataSource,data)
		myStack.currentSize++
	}
}
func (myStack *Stack) IsEmpty() bool {
	if myStack.currentSize == 0 {
		return true
	} else {
		return false
	}
}

// IsFull 判断满了
func (myStack *Stack) IsFull() bool {
	if myStack.currentSize >= myStack.capSize {
		return true
	} else {
		return false
	}
}
