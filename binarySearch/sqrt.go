package binarySearch

func BinarySearchSqrt(target float64) float64 {
	if target < 0 {
		return 0.0
	}
	low := 1.0
	high := target
	if target < 1 {
		low = 0
		high = 1
	}
	mid := (low + high) / 2
	for high-low > 1e-4 {
		mid = (low + high) / 2
		if mid*mid < target {
			low = mid
		} else if mid*mid > target {
			high = mid
		} else {
			break
		}
	}
	return mid
}
