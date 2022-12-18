package main

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want []int
	}{
		{"case-1", nil, nil},
		{"case-2", []int{1}, []int{1}},
		{"case-3", []int{1, 2, 3}, []int{1, 2, 3}},
		{"case-4", []int{3, 2, 1}, []int{1, 2, 3}},
		{"case-5", []int{3, 1, 2}, []int{1, 2, 3}},
		{"case-6", []int{1, 1, 1}, []int{1, 1, 1}},
		{"case-7", []int{2, 0, 5, 7, 12, -1, 23}, []int{-1, 0, 2, 5, 7, 12, 23}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			heapSort(tc.s)
			if !reflect.DeepEqual(tc.s, tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}
