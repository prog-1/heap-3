package main

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{1}, []int{1}},
		{[]int{10, 10}, []int{10, 10}},
		{[]int{-9, 1}, []int{-9, 1}},
		{[]int{20, 5}, []int{5, 20}},
		{[]int{20, 5, 1}, []int{1, 5, 20}},
		{[]int{20, 5, 4, 2, 1}, []int{1, 2, 4, 5}},
		{[]int{5, 20, 4, 2, 1}, []int{1, 2, 4, 5, 20}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 3, 2, 5, 6, 7, 8, 9, 30, 11, 15, 14, 13, 14}, []int{1, 2, 3, 5, 6, 7, 8, 9, 11, 13, 14, 14, 15, 30}},
		{[]int{10, 8, 9, 1, 5}, []int{1, 5, 8, 9, 10}},
	} {
		got := make([]int, len(tc.s))
		copy(got, tc.s)
		heapSort(got)
		if !equal(got, tc.want) {
			t.Errorf("HeapSort(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
