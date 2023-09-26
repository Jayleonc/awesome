package selection

func SorcSelection(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		swap(arr, i, min)
	}
	return arr
}

func swap(arr []int, i int, min int) {
	temp := arr[i]
	arr[i] = arr[min]
	arr[min] = temp
}
