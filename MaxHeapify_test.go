package main

import (
	"reflect"
	"testing"
)

func TestMaxHipify(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input []int
		want  []int
	}{
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"1el", []int{2}, []int{2}},
		{"complete", []int{10, 12, 7, 1, -5, 1, 5, 100, 12}, []int{100, 12, 7, 12, -5, 1, 5, 1, 10}},
		{"full", []int{10, 1, 12, 100, 20, 2, 17}, []int{100, 20, 17, 1, 10, 2, 12}},
		{"1child", []int{10, 1, 12, 100, 20, 2}, []int{100, 20, 12, 1, 10, 2}},
	} {
		if got := MaxHeapify(tc.input); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}
