package Sort

func SelectSortMax(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	} else {
		max := arr[0]
		for i := 0; i < length; i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max
	}

}

func SelectSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ { //只剩一个元素不需要挑选
			min := i //标记索引
			for j := i + 1; j < length; j++ { //每次选出一个极小值
				if arr[min] < arr[j] {
					min = j
				}
			}
		}
	}

}
