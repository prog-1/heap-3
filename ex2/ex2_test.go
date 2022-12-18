package main

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	for _, tc := range []struct {
		name string
		h    []int
		want []int
	}{
		{"1", nil, nil},
		{"2", []int{1}, []int{1}},
		{"3", []int{1, 2, 3}, []int{1, 2, 3}},
		{"4", []int{3, 2, 1}, []int{1, 2, 3}},
		{"5", []int{3, 1, 2}, []int{1, 2, 3}},
		{"6", []int{1, 1, 1}, []int{1, 1, 1}},
		{"7", []int{2, 0, 5, 7, 12, -1, 23}, []int{-1, 0, 2, 5, 7, 12, 23}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			heapSort(tc.h)
			if !reflect.DeepEqual(tc.h, tc.want) {
				t.Errorf("got = %v, want = %v", tc.h, tc.want)
			}
		})
	}
}
