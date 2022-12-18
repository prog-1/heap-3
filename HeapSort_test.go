package main

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input []int
		want  []int
	}{
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"1el", []int{2}, []int{2}},
		{"sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"unsorted", []int{4, 1, 5, 2, 3}, []int{1, 2, 3, 4, 5}},
	} {
		if got := HeapSort(tc.input); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}
