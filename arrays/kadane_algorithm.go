package arrays

// LargestSumContiguousSubarray ... find the sum of contiguous subarray within a one-dimensional array of numbers
// which has the largest sum.
// Time complexity O(n)
func LargestSumContiguousSubarray(input []int) int {
	maxSumSoFar := 0
	localMaxSum := 0

	for i := 0; i < len(input); i++ {
		localMaxSum = localMaxSum + input[i]
		if localMaxSum < 0 {
			localMaxSum = 0
		} else if maxSumSoFar < localMaxSum {
			maxSumSoFar = localMaxSum
		}
	}
	return maxSumSoFar
}
