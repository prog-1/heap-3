package heap_sort

import (
	"heap-3/heapify"
)

func HeapSort(s []int) {
	heapify.Heapify(s)
	for i := len(s) - 1; i >= 0; i-- {
		s[0], s[i] = s[i], s[0]
		heapify.HeapifyUp(0, s[:i])
	}
}
