package leetcode

type Allocator struct {
	n   int
	Arr []int
}

func AllocatorConstructor(n int) Allocator {
	return Allocator{
		n:   n,
		Arr: make([]int, n),
	}
}

func (this *Allocator) Allocate(size int, mID int) int {
	count := 0
	for i := 0; i < len(this.Arr); i++ {
		if this.Arr[i] == 0 {
			count++
		} else {
			count = 0
		}
		if count == size {
			for j := i - size + 1; j < i+1; j++ {
				this.Arr[i] = mID
			}
			return i - size + 1
		}
	}
	return -1
}

func (this *Allocator) FreeMemory(mID int) int {
	numOfmID := 0
	for i := 0; i < len(this.Arr); i++ {
		if this.Arr[i] == mID {
			numOfmID++
			this.Arr[i] = 0
		}
	}
	return numOfmID
}
