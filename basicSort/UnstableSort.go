package basicSort

import "fmt"

func QuickSortInSlice(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
	less := make([]int, 0)
	greater := make([]int, 0)
	for _, v := range arr[1:] {
		if v < pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	return append(append(QuickSortInSlice(less), pivot), QuickSortInSlice(greater)...)
}

// QuickSortInHoare 双指针法交换，函数无需返回，原地排序
func QuickSortInHoare(arr []int, low, high int) {
	if low >= high { // low和high指针用来限定排序的范围。
		return
	}
	pivot := arr[low] // 取范围内的第一个值
	left := low + 1
	right := high

	for left <= right {
		// 左指针找 >= pivot 的元素
		for left <= high && arr[left] < pivot { // 应限定左指针的遍历范围，必须小于high
			left++ // 没找到就一直右移
		}
		// 右指针找 <= pivot 的元素
		for right > low && arr[right] > pivot { // 同理，右指针大于low
			right-- // 没找到就一直左移
		}
		if left <= right { // 找到了且没有互换位置
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	// 将基准交换到正确位置（右指针最终位置）
	arr[low], arr[right] = arr[right], arr[low]
	fmt.Println(arr)
	QuickSortInHoare(arr, low, right-1)
	QuickSortInHoare(arr, right+1, high)
}

// SelectionSort 思路：每轮循环中找到最小值放到已经排序的位置
func SelectionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	sortedPosition := 0
	for sortedPosition < len(arr) {
		curMinIdx := sortedPosition
		curMinVal := arr[curMinIdx]
		for i := sortedPosition; i < len(arr); i++ {
			if arr[i] < curMinVal {
				curMinIdx = i
				curMinVal = arr[i]
			}
		}
		arr[sortedPosition], arr[curMinIdx] = arr[curMinIdx], arr[sortedPosition]
		sortedPosition++
	}
	return arr
}

type MaxHeap struct {
	data []int
	size int
}

// Heapify 堆化过程
func (heap *MaxHeap) Heapify() {
	
}

func MaxHeapConstructor() *MaxHeap {

}

func HeapSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

}
