package main

import (
	"./ArrayList"
	"./CricleQueue"
	"./Queue"
	"fmt"
)

func main1() {
	list := ArrayList.NewArrayList()
	list.Append("2")
	list.Append("b")
	list.Append("c")
	for i := 0; i < 20; i++ {
		list.Append(i)
	}
	fmt.Println(list)
}

func main2() {
	list := ArrayList.NewArrayList()
	list.Append("2")
	list.Append("b")
	list.Append("c")
	for i := 0; i < 10; i++ {
		fmt.Println(list)
		list.Insert(1, "Q")
	}
	list.Delete(1)
	fmt.Println(list)
}

func mai3() {
	list := ArrayList.NewArrayList()
	list.Append("2")
	list.Append("b")
	list.Append("c")
	for it := list.Iterator(); it.HasNext(); {
		item, _ := it.Next()
		fmt.Println(item)
	}
}

func main4() {
	mystack := ArrayList.NewArrayListStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	mystack.Push(4)
	mystack.Push(5)
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
}

func Add(num int) int {
	if num == 0 {
		return 0
	} else {
		return num + Add(num-1)
	}
}

func main6() {
	mystack := ArrayList.NewArrayListStackX()
	mystack.PushX(1)
	mystack.PushX(2)
	mystack.PushX(3)
	mystack.PushX(4)
	mystack.PushX(5)
	for it := mystack.MyIt; it.HasNext(); {
		item, _ := it.Next()
		fmt.Println(item)
	}
	//fmt.Println(Add(5))
}

// 栈模拟递归
func main7() {
	mystack := ArrayList.NewArrayListStack()
	mystack.Push(5)
	last := 0
	for !mystack.IsEmpty() {
		data := mystack.Pop()
		if data == 0 {
			last += 0
		} else {
			last += data.(int)
			mystack.Push((data.(int) - 1))
		}

	}
	fmt.Println(last)
}

func FAB(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else {
		return FAB(num-1) + FAB(num-2)
	}
}

// 栈模拟递归--斐波那契数列
func main45() {
	fmt.Println(FAB(7))
	mystack := ArrayList.NewArrayListStack()
	mystack.Push(7)
	last := 0
	for !mystack.IsEmpty() {
		data := mystack.Pop()
		if data == 1 || data == 2 {
			last += 1
		} else {
			mystack.Push((data.(int) - 2))
			mystack.Push((data.(int) - 1))
		}

	}
	fmt.Println(last)
}

// 队列使用（队列是广度遍历，栈是深度遍历）
func main12() {
	queue := Queue.NewQueue()
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.Size())
	queue.EnQueue("sb")
	queue.EnQueue("sb2")
	queue.EnQueue("sb3")
	queue.EnQueue("sb4")
	fmt.Println(queue.Size())
	fmt.Println(queue.Front())
	fmt.Println(queue.End())
	fmt.Println(queue.Size())
	queue.DeQueue()
	fmt.Println(queue.Size())
	queue.DeQueue()
	fmt.Println(queue.Size())
	queue.DeQueue()
	fmt.Println(queue.Size())
	queue.DeQueue()
	fmt.Println(queue.Size())
}

func main() {
	var q CricleQueue.CircleQueue
	CricleQueue.InitQueue(&q)
	CricleQueue.EnQueue(&q,1)
	CricleQueue.EnQueue(&q,2)
	CricleQueue.EnQueue(&q,3)
	CricleQueue.EnQueue(&q,4)
	fmt.Println(CricleQueue.DeQueue(&q))
	fmt.Println(CricleQueue.DeQueue(&q))
	fmt.Println(CricleQueue.DeQueue(&q))
	fmt.Println(CricleQueue.DeQueue(&q))

}
