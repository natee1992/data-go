package main

import (
	"./Link"
	"fmt"
)

func mainxs() {

	node1 := new(Link.Node)
	node2 := new(Link.Node)
	node3 := new(Link.Node)
	node4 := new(Link.Node)
	node1.Data = 1
	node1.PNext = node2
	node2.Data = 2
	node2.PNext = node3
	node3.Data = 3
	node3.PNext = node3
	node4.Data = 4
	fmt.Println(node1.PNext.Data)
}

/*
栈，先进后出，操作口只有一个，头部删除头部插入，尾部插入，尾部删除
队列 头部插入，尾部删除，尾部插入，头部删除
*/

func mainxs1() {
	myStack := Link.NewStack()
	for i := 0; i < 100; i++ {
		myStack.Push(i)
	}
	for data := myStack.Pop(); data != nil; data = myStack.Pop() {
		fmt.Println(data)
	}
}

func main10() {
	myStack := Link.NewLinkQueue()
	for i := 0; i < 10; i++ {
		myStack.Enqueue(i)
	}
	for data := myStack.Dequeue(); data != nil; data = myStack.Dequeue() {
		fmt.Println(data)
	}
}
