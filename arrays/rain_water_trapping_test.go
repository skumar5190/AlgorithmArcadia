package arrays

import "testing"

var benchmarkInputRainWater []int


func TestGetWaterTrapped(t *testing.T) {
	cases := []struct {
		name           string
		input          []int
		expectedOutput int
	}{
		{
			name:           "should give expected output",
			input:          []int{3, 0, 2, 0, 4},
			expectedOutput: 7,
		},
		{
			name:           "should give zero output",
			input:          []int{1, 2},
			expectedOutput: 0,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			result := GetMaxWaterTrapped(c.input)
			if result != c.expectedOutput {
				t.Errorf("incorrect result: actual-%v , epected-%v", result, c.expectedOutput)
			}
		})
	}
}

func TestGetWaterTrappedOptimized(t *testing.T) {
	cases := []struct {
		name           string
		input          []int
		expectedOutput int
	}{
		{
			name:           "should give expected output",
			input:          []int{3, 0, 2, 0, 4},
			expectedOutput: 7,
		},
		{
			name:           "should give zero output",
			input:          []int{1, 2},
			expectedOutput: 0,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			result := GetMaxWaterTrappedOptimized(c.input)
			if result != c.expectedOutput {
				t.Errorf("incorrect result: actual-%v , epected-%v", result, c.expectedOutput)
			}
		})
	}
}

func BenchmarkGetWaterTrapped(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMaxWaterTrapped(benchmarkInput)
	}
}

func BenchmarkGetWaterTrappedOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMaxWaterTrappedOptimized(benchmarkInput)
	}
}

func init() {
	benchmarkInput = []int{7, 1, 2, 3, 4, 5, 3, 2, 3, 7, 8, 9, 1, 2, 3, 7, 9, 1, 2, 5, 9, 10,
		7, 1, 2, 3, 4, 5, 3, 2, 3, 7, 8, 9, 1, 2, 3, 7, 9, 1, 2, 5, 9, 10,
		7, 1, 2, 3, 4, 5, 3, 2, 3, 7, 8, 9, 1, 2, 3, 7, 9, 1, 2, 5, 9, 10,
		7, 1, 2, 3, 4, 5, 3, 2, 3, 7, 8, 9, 1, 2, 3, 7, 9, 1, 2, 5, 9, 10,
		7, 1, 2, 3, 4, 5, 3, 2, 3, 7, 8, 9, 1, 2, 3, 7, 9, 1, 2, 5, 9, 10}
}
