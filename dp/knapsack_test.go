package dp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZeroOneKnapsack(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		weights  []int
		capacity int
		expected int
	}{
		{
			name:     "case1",
			values:   []int{8, 6, 7, 5},
			weights:  []int{2, 1, 2, 1},
			capacity: 4,
			expected: 19,
		},
		{
			name:     "case2 - zero capacity",
			values:   []int{8, 6, 7, 5},
			weights:  []int{2, 1, 2, 1},
			capacity: 0,
			expected: 0,
		},
		{
			name:     "case3 - insufficient capacity",
			values:   []int{10, 20, 30},
			weights:  []int{1, 3, 4},
			capacity: 2,
			expected: 10,
		},
		{
			name:     "case4 - large capacity",
			values:   []int{10, 20, 30},
			weights:  []int{1, 3, 4},
			capacity: 8,
			expected: 60,
		},
	}
	// 遍历每个测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ZeroOneKnapsack(tt.capacity, tt.values, tt.weights)
			assert.Equal(t, tt.expected, result, "Test case %s failed", tt.name)
		})
	}
}
