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

func main() {
	list := ArrayList.NewArrayList()
	list.Append("2")
	list.Append("b")
	list.Append("c")
	for it := list.Iterator(); it.HasNext(); {
		item, _ := it.Next()
		fmt.Println(item)
	}
}
