package main

import (
	"./ArrayList"
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

func main() {
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
}
