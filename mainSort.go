package main

import (
	"./Sort"
	"fmt"
)

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 0}
	fmt.Println(Sort.SelectSortMax(arr))
}
