package Sort

import "fmt"

//银行，每人都有存款
// 100 1000 10000

//800万高考考生, 分数排序 0-750 1500
//1亿人，身高排序 0-300

func RadixSort(arr []int) []int {
	max := SelectSortMax(arr) //寻找数组极大值
	for bit := 1; max/bit > 0; bit *= 10 {
		//按照数量级分段
		fmt.Println(arr)
		arr = BitSort(arr, bit) //每次处理一个级别的排序
	}
	return arr
}

func BitSort(arr []int, bit int) []int {
	length := len(arr)           //数组长度
	bitCounts := make([]int, 10) //统计长度
	for i := 0; i < length; i++ {
		num := (arr[i] / bit) % 10
		bitCounts[num]++ //统计余数相等的个数
	}
	fmt.Println(bitCounts)
	// 0 1 2 3 4 5
	// 1 0 3 0 0 1
	for i := 1; i < 10; i++ {
		bitCounts[i] += bitCounts[i-1] //叠加，计算位置分布
	}
	fmt.Println(bitCounts)
	tmp := make([]int, 10)
	for i := length - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10     //分层
		tmp[bitCounts[num]-1] = arr[i] //计算排序的位置
		bitCounts[num]--
	}
	for i := 0; i < length; i++ {
		arr[i] = tmp[i]
	}
	return arr
}
