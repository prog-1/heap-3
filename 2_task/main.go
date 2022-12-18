package main

import "fmt"

func heapify(s []int) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		heapifyUp(s, i)
	}
}

func heapifyUp(s []int, i int) {
	l := i*2 + 1
	r := i*2 + 2
	if r < len(s) && s[r] > s[l] && l < len(s) {
		l = r
	}
	if l != i {
		s[i], s[l] = s[l], s[i]
		heapifyUp(s, l)
	}
}

func heapSort(s []int) {
	heapify(s)
	for i := len(s) - 1; i > 0; i-- {
		s[0], s[i] = s[i], s[0]
		heapifyUp(s[:i], 0)
	}
}

func main() {
	s := []int{1, 3, 2, 5, 6, 7, 8, 9, 30, 11, 15, 14, 13, 14}
	heapSort(s)
	fmt.Println(s)
}
