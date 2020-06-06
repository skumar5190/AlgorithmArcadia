package matrix

import "testing"

var benchmarkMatrix [][]int

func init() {
	benchmarkMatrix = [][]int{
		{1, 1, 0, 0, 0, 0, 1, 1, 0, 1, 0, 1, 0, 1},
		{0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1},
		{1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{1, 0, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 0},
	}
}

func TestNumberOfIsland(t *testing.T) {
	cases := []struct {
		name           string
		input          [][]int
		expectedOutput int
	}{
		{
			name: "should give expected output",
			input: [][]int{
				{1, 1, 0, 0, 0},
				{0, 1, 0, 0, 1},
				{1, 0, 0, 1, 1},
				{0, 0, 0, 0, 0},
				{1, 0, 1, 0, 1},
			},
			expectedOutput: 5,
		},
		{
			name: "should give expected output when one's are connected",
			input: [][]int{
				{1, 1, 0, 0, 0},
				{0, 1, 0, 0, 1},
				{1, 1, 1, 1, 1},
				{0, 0, 1, 0, 0},
				{1, 1, 0, 1, 1},
			},
			expectedOutput: 1,
		},
		{
			name:           "should return zero when matrix is empty",
			input:          [][]int{},
			expectedOutput: 0,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			result := NumberOfIsland(c.input)
			if result != c.expectedOutput {
				t.Errorf("incorrect result: actual-%v , epected-%v", result, c.expectedOutput)
			}
		})
	}
}

func BenchmarkNumberOfIsland(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumberOfIsland(benchmarkMatrix)
	}
}
