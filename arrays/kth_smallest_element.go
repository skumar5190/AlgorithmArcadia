package arrays

import (
	"container/heap"
	"sort"
)

func GetKthSmallestElementUsingSort(input []int, k int) int {
	if len(input) < k {
		return -1
	}
	sort.Ints(input)
	return input[k-1]
}

func GetKthSmallestElementUsingMaxHeap(input []int, k int) int {
	if len(input) < k {
		return -1
	}
	h := make(IntMaxHeap, 0, k)
	for i := 0; i < k; i++ {
		h = append(h, input[i])
	}
	heap.Init(&h)

	for i := k; i < len(input); i++ {
		if h[0] > input[i] {
			heap.Pop(&h)
			heap.Push(&h, input[i])
		}
	}
	return heap.Pop(&h).(int)
}

func GetKthSmallestElementUsingMinHeap(input []int, k int) int {
	if len(input) < k {
		return -1
	}
	h := make(IntMinHeap, 0, len(input))
	for i := 0; i < len(input); i++ {
		h = append(h, input[i])
	}
	heap.Init(&h)
	for i := 1; i < k; i++ {
		heap.Pop(&h)
	}
	return heap.Pop(&h).(int)
}
