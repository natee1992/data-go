package Sort

func ShellSortStep(arr []int, start int, gap int) {
	length := len(arr)
	for i := start+gap;i<length;i+=gap{ //插入排序的变种
		backup := arr[i] //备份插入的数据
		j := i - gap
		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j] //从后往前移动
			j-=gap
		}
		arr[j+gap] = backup
		//fmt.Println(arr)
	}
}

// 多线程排序
func ShellSort(arr []int) []int {
	length := len(arr)
	if length < 1 {
		return arr
	} else {
		gap := length / 2
		for gap > 0 {
			for i := 0; i < gap; i++ { //处理每个元素的步长
				ShellSortStep(arr, i, gap)
			}
			gap /= 2

		}
	}
	return arr
}
