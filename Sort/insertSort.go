package Sort

// InsertSort: 插入排序
func InsertSort(arr []int) []int {
	backup := arr[2]
	j := 2 - 1
	for j >= 0 && backup < arr[j] {
		arr[j+1] = arr[j] //从前往后移动
		j--
	}

	arr[j+1] = backup //插入
	return arr
}

func InsertSortAll(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	} else {
		for i := 1; i < len(arr); i++ { //跳过第一个
			backup := arr[i] //备份插入的数据
			j := i - 1
			for j >= 0 && backup < arr[j] {
				arr[j+1] = arr[j] //从后往前移动
				j--
			}
			arr[j+1] = backup
		}

	}
	return arr
}
