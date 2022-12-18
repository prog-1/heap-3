package main

import (
	"testing"
)

func TestHeapify(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{10, 10}, []int{10, 10}},
		{[]int{-9, 1}, []int{1, -9}},
		{[]int{20, 5}, []int{20, 5}},
		{[]int{20, 5, 1}, []int{20, 5, 1}},
		{[]int{20, 5, 4, 2, 1}, []int{20, 5, 4, 2, 1}},
		{[]int{5, 20, 4, 2, 1}, []int{20, 5, 4, 2, 1}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 1, 2}},
		{[]int{1, 3, 2, 5, 6, 7, 8, 9, 30, 11, 15, 14, 13, 14}, []int{30, 15, 14, 9, 11, 13, 14, 3, 5, 1, 6, 7, 2, 8}},
		{[]int{1, 5, 8, 9, 10}, []int{10, 9, 8, 1, 5}},
	} {
		got := make([]int, len(tc.s))
		copy(got, tc.s)
		heapify(got)
		if !equal(got, tc.want) {
			t.Errorf("Heapify(%v) = %v, want = %v", tc.s, got, tc.want)
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
