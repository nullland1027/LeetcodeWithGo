package basicSort

func BubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	isSorted := true // 标志位，如果某一轮没有交换，表示数组已经有序，可以提前终止
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				isSorted = false
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if isSorted {
			break
		}
	}
	return arr
}

func insertInSortedSlice(slice []int, value int) []int {
	if len(slice) == 0 {
		return []int{value}
	}
	if value <= slice[0] {
		return append([]int{value}, slice...)
	}
	if value > slice[len(slice)-1] {
		return append(slice, value)
	}
	for i := 0; i < len(slice); i++ {
		if value <= slice[i] {
			return append(slice[:i], append([]int{value}, slice[i:]...)...)
		}
	}
	return slice
}

func InsertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// 默认[0]位置已经有序
	sortedSlice := []int{arr[0]}
	for i := 1; i < len(arr); i++ {
		sortedSlice = insertInSortedSlice(sortedSlice, arr[i])
	}
	return sortedSlice
}

func mergeTwoSortedSlice(arr1, arr2 []int) []int {
	ans := make([]int, 0)
	p1, p2 := 0, 0
	for p1 < len(arr1) && p2 < len(arr2) {
		if arr1[p1] <= arr2[p2] {
			ans = append(ans, arr1[p1])
			p1++
		} else {
			ans = append(ans, arr2[p2])
			p2++
		}
	}
	if p1 < len(arr1) {
		ans = append(ans, arr1[p1:]...)
	}
	if p2 < len(arr2) {
		ans = append(ans, arr2[p2:]...)
	}
	return ans
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// Divide and conquer
	midIdx := len(arr) / 2
	leftPart := arr[:midIdx]
	rightPart := arr[midIdx:]
	sortedLeft := MergeSort(leftPart)
	sortedRight := MergeSort(rightPart)
	return mergeTwoSortedSlice(sortedLeft, sortedRight)
}
