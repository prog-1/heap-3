# Exercise 1

Implement a function (with tests) `func heapify(s []int)` in the directory
`heapify` that transforms a given slice into a max heap.

> Hint: As a helper function, implement `heapifyUp(s []int, i int)` function
that swaps elements in the heap `s` (the heap property may not hold for an
element with index `i` in the heap), so that `i` is inserted into correct
position.

# Exercise 2

Implement a function (with tests) `func heapSort(s []int)` in the directory `heap_sort` that implements a heap sort algorithm.

> Hint: Call the function `heapifyUp(s, 0)`.
