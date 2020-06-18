package Sort

import (
	"strings"
)

// SelectSort： 选择排序
func SelectSortStringMax(arr []string) string {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	} else {
		max := arr[0]
		for i := 0; i < length-1; i++ { //只剩一个元素不需要挑选
			if strings.Compare(arr[i],max) < 0{
				max = arr[i]
			}
		}
		return max
	}

}


// SelectSort： 选择排序
func SelectSortString(arr []string) []string {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ { //只剩一个元素不需要挑选
			min := i //标记索引
			for j := i + 1; j < length; j++ { //每次选出一个极小值
				if strings.Compare(arr[min],arr[j]) > 0 {
					min = j //保存极小值的索引
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i] // 数据交换
			}
		}
	}
	return arr

}
