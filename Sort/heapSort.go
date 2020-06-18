package Sort

// heapSortMax: 堆排序找出最大值
func HeapSortMax(arr [] int,length int) []int {
	if length == 1 {
		return arr
	} else {
		depth := length/2 - 1 //深度 n ,2*n+1 ,2*n+2
		for i := depth; i >= 0; i-- { //循环所有的三节点
			topMax := i //假定最大的在i的位置
			leftChild := 2*i + 1
			rightChild := 2*i + 2
			if leftChild <= length-1 && arr[leftChild] > arr[topMax] {
				topMax = leftChild
			}
			if rightChild <= length-1 && arr[rightChild] > arr[topMax] {
				topMax = rightChild
			}
			if topMax != i { //确保i值最大
				arr[i], arr[topMax] = arr[topMax], arr[i]
			}
		}
		return arr
	}
}

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastmesslen := length - i //每次截取一段
		HeapSortMax(arr, lastmesslen)
		if i < length {
			arr[0], arr[lastmesslen-1] = arr[lastmesslen-1], arr[0]
		}
	}
	return arr
}
