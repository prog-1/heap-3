package main

import (
	"reflect"
	"testing"
)

func TestHeapify(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want []int
	}{
		{"test 1", []int{7, 9, 2, 15, 10, 5, 12}, []int{15, 10, 12, 9, 7, 5, 2}},
		{"test 2", []int{4, 6, 2, 7, 1, 5, 3}, []int{7, 6, 5, 4, 1, 2, 3}},
		{"test 3", []int{1, 14, 7, 8, 3, 9, 5}, []int{14, 8, 9, 1, 3, 7, 5}},
		{"with negative", []int{-4, 6, 2, -7, 1, -5, 3}, []int{6, 1, 3, -7, -4, -5, 2}},
		{"with 0", []int{4, 0, 2, 7, 0, 5, 3}, []int{7, 4, 5, 0, 0, 2, 3}},
		{"one el", []int{7}, []int{7}},
		{"empty", []int{}, []int{}},
		{"nil", nil, nil},
		//reliability of results is verified by me
	} {
		t.Run(tc.name, func(t *testing.T) {
			heapify(tc.s)
			if !reflect.DeepEqual(tc.s, tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}
