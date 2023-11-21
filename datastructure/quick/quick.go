package quick

func Partition(a []int, low, high int) int {
	pos := a[low]
	for low < high {
		// 判断右边是不是大于基准值，如果是就不用动，high--
		// 否则，把high值放到low
		for low < high && a[high] >= pos {
			high--
		}
		a[low] = a[high] // high 比基准值小，放到左边

		// 判断左边是不是小于基准值，如果是就不用动（小的就应该在左边），low++
		for low < high && a[low] <= pos {
			low++
		}
		a[high] = a[low] // low 比基准值大，放到右边

	}
	// 当 low 不等于 high，认为两者相遇，返回当前基准元素的最终位置
	a[low] = pos
	return low
}

func Sort(a []int, low, high int) []int {
	if low < high {
		pivotPos := Partition(a, low, high)
		Sort(a, low, pivotPos-1)
		Sort(a, pivotPos+1, high)
	}
	return a
}
