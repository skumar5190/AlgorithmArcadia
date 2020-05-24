package arrays

import (
	"testing"
)

var benchmarkInput []int

func init() {
	benchmarkInput = createInputForBenchmark()
}

func TestGetKthSmallestElementUsingSort(t *testing.T) {
	cases := []struct {
		name           string
		input          []int
		k              int
		expectedOutput int
	}{
		{
			name:           "should give expected output",
			input:          []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			k:              3,
			expectedOutput: 3,
		},
		{
			name:           "should give minus one output",
			input:          []int{1, 2},
			k:              3,
			expectedOutput: -1,
		},
		{
			name:           "should give expected  output array is sorted in descending order",
			input:          []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			k:              4,
			expectedOutput: 4,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			result := GetKthSmallestElementUsingSort(c.input, c.k)
			if result != c.expectedOutput {
				t.Errorf("incorrect result: actual-%v , epected-%v", result, c.expectedOutput)
			}
		})
	}
}

func TestGetKthSmallestElementUsingMaxHeap(t *testing.T) {
	cases := []struct {
		name           string
		input          []int
		k              int
		expectedOutput int
	}{
		{
			name:           "should give expected output",
			input:          []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			k:              3,
			expectedOutput: 3,
		},
		{
			name:           "should give minus one output",
			input:          []int{1, 2},
			k:              3,
			expectedOutput: -1,
		},
		{
			name:           "should give expected  output array is sorted in descending order",
			input:          []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			k:              4,
			expectedOutput: 4,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			result := GetKthSmallestElementUsingMaxHeap(c.input, c.k)
			if result != c.expectedOutput {
				t.Errorf("incorrect result: actual-%v , epected-%v", result, c.expectedOutput)
			}
		})
	}
}

func TestGetKthSmallestElementUsingMinHeap(t *testing.T) {
	cases := []struct {
		name           string
		input          []int
		k              int
		expectedOutput int
	}{
		{
			name:           "should give expected output",
			input:          []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			k:              3,
			expectedOutput: 3,
		},
		{
			name:           "should give minus one output",
			input:          []int{1, 2},
			k:              3,
			expectedOutput: -1,
		},
		{
			name:           "should give expected  output array is sorted in descending order",
			input:          []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			k:              4,
			expectedOutput: 4,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			result := GetKthSmallestElementUsingMinHeap(c.input, c.k)
			if result != c.expectedOutput {
				t.Errorf("incorrect result: actual-%v , epected-%v", result, c.expectedOutput)
			}
		})
	}
}

func BenchmarkGetKthSmallestElementUsingMaxHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetKthSmallestElementUsingMaxHeap(benchmarkInput, 10)
	}
}

func BenchmarkGetKthSmallestElementUsingSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetKthSmallestElementUsingSort(benchmarkInput, 10)
	}
}

func BenchmarkGetKthSmallestElementUsingMinHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetKthSmallestElementUsingMinHeap(benchmarkInput, 10)
	}
}

func createInputForBenchmark() []int {

	input := make([]int, 0, 100)
	for i := 1; i <= 100; i++ {
		input = append(input, i)
	}
	return input
}
