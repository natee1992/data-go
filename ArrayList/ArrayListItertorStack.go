package ArrayList

type StackArrayX interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
}

type StackX struct {
	myArray     *ArrayList
	MyIt	Iterator
}

func NewArrayListStackX() *StackX {
	myStack := new(StackX)
	myStack.myArray = NewArrayList()
	myStack.MyIt = myStack.myArray.Iterator()
	return myStack
}

func (myStack *StackX) ClearX() {
	myStack.myArray.Clear()
	myStack.myArray.theSize = 0
}

func (myStack *StackX) SizeX() int {
	return myStack.myArray.Size()
}

func (myStack *StackX) PopX() interface{} {
	if !myStack.IsEmptyX() {
		last := myStack.myArray.dataStore[myStack.myArray.theSize-1]
		myStack.myArray.Delete(myStack.myArray.theSize - 1)
		return last
	}
	return nil
}
func (myStack *StackX) PushX(data interface{}) {
	if !myStack.IsFullX() {
		myStack.myArray.Append(data)
	}
}
func (myStack *StackX) IsEmptyX() bool {
	if myStack.myArray.theSize == 0 {
		return true
	} else {
		return false
	}
}

// IsFull 判断满了
func (myStack *StackX) IsFullX() bool {
	if myStack.myArray.theSize >= 10 {
		return true
	} else {
		return false
	}
}
