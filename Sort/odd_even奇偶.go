package Sort

// 1 7 3 8 2 9 5 6 4 0

// 1 3 2 5 4
// 7 8 9 6 0

func OddEven(arr []int) []int {
	isSorted := false
	for ; isSorted == false; {
		isSorted = true
		for i := 1; i < len(arr)-1; i = i + 2 { //奇数位
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
		for i := 0; i < len(arr)-1; i = i + 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
	}
	return arr
}
