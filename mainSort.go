package main

import (
	"./Sort"
	"fmt"
)

func main01() {
	arr := []int{1021, 912, 22, 83, 312312, 721, 411, 61, 5, 0}
	//fmt.Println(Sort.SelectSort(arr))
	//fmt.Println(Sort.BubbleSort(arr))
	//fmt.Println(Sort.InsertSort(arr))
	//fmt.Println(Sort.InsertSortAll(arr))
	//fmt.Println(Sort.HeapSortMax(arr,10))
	//fmt.Println(Sort.HeapSort(arr))
	//fmt.Println(Sort.OddEven(arr))
	//fmt.Println(Sort.MergeSort(arr))
	//fmt.Println(Sort.ShellSort(arr))
	fmt.Println(Sort.RadixSort(arr))
}

func main0() {
	//fmt.Println(strings.Compare("a", "b"))
	//fmt.Println(strings.Compare("a1", "a2"))
	//pb := "b"
	//pa := "a"
	//fmt.Println(&pa, &pb)
	ss := []string{"a", "c", "cx", "bc", "ba"}
	fmt.Println(Sort.SelectSortString(ss))
}

func main()  {
	//Sort.SortCsdn("/Users/natee/Downloads/爱租房+微信公众号tornado/day2/CSDN-中文IT社区-600万.sql")
	Sort.GetSortTime("CSDNpass.txt")
}