package main

import (
	"reflect"
	"testing"
)

func TestHeapsort(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want []int
	}{
		{"test 1", []int{7, 9, 2, 15, 10, 5, 12}, []int{2, 5, 7, 9, 10, 12, 15}},
		{"test 2", []int{4, 6, 2, 7, 1, 5, 3}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"test 3", []int{1, 14, 7, 8, 3, 9, 5}, []int{1, 3, 5, 7, 8, 9, 14}},
		{"with negative", []int{-4, 6, 2, -7, 1, -5, 3}, []int{-7, -5, -4, 1, 2, 3, 6}},
		{"with 0", []int{4, 0, 2, 7, 0, 5, 3}, []int{0, 0, 2, 3, 4, 5, 7}},
		{"one el", []int{7}, []int{7}},
		{"empty", []int{}, []int{}},
		{"nil", nil, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			heapsort(tc.s)
			if !reflect.DeepEqual(tc.s, tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}
