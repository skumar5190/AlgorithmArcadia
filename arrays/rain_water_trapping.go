package arrays

// GetMaxWaterTrapped ... will return the water trapped inside a histogram which is represented by barsay of bars.
// Time complexity O(n2)
// Space complexity O(1)
func GetMaxWaterTrapped(bars []int) int {
	result := 0

	// For every element of the barsay
	for i := 1; i < len(bars)-1; i++ {

		// Find the maximum element on its left
		left := bars[i]
		for j := 0; j < i; j++ {
			left = Max(left, bars[j])
		}

		// Find the maximum element on its right
		right := bars[i]
		for j := i + 1; j < len(bars); j++ {
			right = Max(right, bars[j])
		}

		// Update the maximum water
		result = result + (Min(left, right) - bars[i])
	}

	return result
}

// GetMaxWaterTrappedOptimized ... will return the water trapped inside a histogram which is represented by barsay of bars.
// Time complexity O(n)
// Space complexity O(n)
func GetMaxWaterTrappedOptimized(bars []int) int {

	length := len(bars)
	leftMax := make([]int, length, length)
	rightMax := make([]int, length, length)

	result := 0

	leftMax[0] = bars[0]
	for i := 1; i < length; i++ {
		leftMax[i] = Max(leftMax[i-1], bars[i])
	}
	rightMax[length-1] = bars[length-1]
	for i := length - 2; i >= 0; i-- {
		rightMax[i] = Max(rightMax[i+1], bars[i])
	}

	for i := 0; i < length; i++ {
		result += Min(leftMax[i], rightMax[i]) - bars[i]
	}

	return result
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Max returns the larger of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
