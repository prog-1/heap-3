package main

import (
	"reflect"
	"testing"
)

func TestHeapify(t *testing.T) {
	for _, tc := range []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{10, 9, 7, 8, 5, 6, 3, 1, 4, 2}},
		{[]int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}, []int{5, 5, 4, 4, 2, 3, 1, 3, 2, 1}},
	} {
		got := make([]int, len(tc.input))
		copy(got, tc.input)
		if heapify(got); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("heapify(%v) = %v, want = %v", tc.input, got, tc.want)
		}
	}
}
