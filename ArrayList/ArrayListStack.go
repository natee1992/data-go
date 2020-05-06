package ArrayList

type StackArray interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
}

type Stack struct {
	myArray     *ArrayList
	capSize     int // 最大范围
	currentSize int //当前实际使用大小
}

func NewArrayListStack() *Stack {
	myStack := new(Stack)
	myStack.myArray = NewArrayList()
	myStack.capSize = 10
	return myStack
}

func (myStack *Stack) Clear() {
	myStack.myArray.Clear()
	myStack.capSize = 10
}

func (myStack *Stack) Size() int {
	return myStack.myArray.Size()
}

func (myStack *Stack) Pop() interface{} {
	if !myStack.IsEmpty() {
		last := myStack.myArray.dataStore[myStack.myArray.theSize-1]
		myStack.myArray.Delete(myStack.myArray.theSize - 1)
		return last
	}
	return nil
}
func (myStack *Stack) Push(data interface{}) {
	if !myStack.IsFull() {
		myStack.myArray.Append(data)
	}
}
func (myStack *Stack) IsEmpty() bool {
	if myStack.myArray.theSize == 0 {
		return true
	} else {
		return false
	}
}

// IsFull 判断满了
func (myStack *Stack) IsFull() bool {
	if myStack.myArray.theSize >= myStack.capSize {
		return true
	} else {
		return false
	}
}
