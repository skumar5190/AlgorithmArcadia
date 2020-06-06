package arrays

import (
	"testing"
)

var benchmarkInputKadaneAlgo []int

func init() {
	benchmarkInputKadaneAlgo := make([]int, 0, 1000)
	for i := 1; i <= 1000; i++ {
		benchmarkInputKadaneAlgo = append(benchmarkInputKadaneAlgo, i)
	}
}

func TestLargestSumContiguousSubarray(t *testing.T) {
	cases := []struct {
		name           string
		input          []int
		expectedOutput int
	}{
		{
			name:           "should give expected output",
			input:          []int{-2, -3, 4, -1, -2, 1, 5, -3},
			expectedOutput: 7,
		},
		{
			name:           "should give expected output when all elements are positive",
			input:          []int{1,2,3,4,5},
			expectedOutput: 15,
		},
		{
			name:           "should give zero if all elements are negative",
			input:          []int{-1,-2,-3,-4},
			expectedOutput: 0,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			result := LargestSumContiguousSubarray(c.input)
			if result != c.expectedOutput {
				t.Errorf("incorrect result: actual-%v , epected-%v", result, c.expectedOutput)
			}
		})
	}
}

func BenchmarkLargestSumContiguousSubarray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestSumContiguousSubarray(benchmarkInputKadaneAlgo)
	}
}

