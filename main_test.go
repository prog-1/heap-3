package main

import (
	"reflect"
	"testing"
)

func TestHipify(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input []int
		want  []int
	}{
		{"1", nil, nil},
		{"2", []int{}, []int{}},
		{"3", []int{2}, []int{2}},
		{"4", []int{10, 12, 7, 1, -5, 1, 5, 100, 12}, []int{100, 12, 7, 12, -5, 1, 5, 1, 10}},
		{"5", []int{10, 1, 12, 100, 20, 2, 17}, []int{100, 20, 17, 1, 10, 2, 12}},
		{"6", []int{10, 1, 12, 100, 20, 2}, []int{100, 20, 12, 1, 10, 2}},
	} {
		if got := heapify(tc.input); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}

func TestHeapSort(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
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
			heapSort(tc.s)
			if !reflect.DeepEqual(tc.s, tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}
