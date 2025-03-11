package leetcode

type OrderedStream struct {
	Value    []string
	ptr      int
	IsFilled []bool
}

func Constructor(n int) OrderedStream {
	os := OrderedStream{
		Value:    make([]string, n+1),
		ptr:      1,
		IsFilled: make([]bool, n+1),
	}
	for i := 1; i <= n; i++ {
		os.IsFilled[i] = false
	}
	return os
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.Value[idKey] = value
	this.IsFilled[idKey] = true
	if !this.IsFilled[idKey] {
		return []string{}
	}
	res := make([]string, 0)
	for i := this.ptr; i < len(this.Value); i++ {
		if this.IsFilled[i] {
			res = append(res, this.Value[i])
			this.ptr++
		} else {
			break
		}
	}
	return res
}
