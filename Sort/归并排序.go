package Sort

func merge(leftarr []string, rightarr []string) []string {
	leftindex := 0
	rightindex := 0
	lastarr := []string{}
	for leftindex < len(leftarr) && rightindex < len(rightarr) {
		if leftarr[leftindex] < rightarr[rightindex] {
			lastarr = append(lastarr, leftarr[leftindex])
			leftindex++
		} else if leftarr[leftindex] > rightarr[rightindex] {
			lastarr = append(lastarr, rightarr[rightindex])
			rightindex++
		} else {
			lastarr = append(lastarr, rightarr[rightindex])
			lastarr = append(lastarr, leftarr[leftindex])
			leftindex++
			rightindex++
		}
	}
	for leftindex < len(leftarr) {
		lastarr = append(lastarr, leftarr[leftindex])
		leftindex++
	}
	for rightindex < len(rightarr) {
		lastarr = append(lastarr, rightarr[rightindex])
		rightindex++
	}
	return lastarr
}

func MergeSort(arr []string) []string {
	lenght := len(arr)
	if lenght <= 1 {
		return arr
	} else {
		mid := lenght / 2
		leftarr := MergeSort(arr[:mid])
		rightarr := MergeSort(arr[mid:])
		return merge(leftarr, rightarr)
	}
}
