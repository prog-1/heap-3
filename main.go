package main

import (
	"fmt"
	"heap-3/heap_sort"
	"heap-3/heapify"
)

func main() {
	s := []int{6, 0, 1, 5, 9, 4, 2, 3, 7, 8}
	fmt.Println(s)
	heapify.Heapify(s)
	fmt.Println(s)
	heap_sort.HeapSort(s)
	fmt.Println(s)
}
