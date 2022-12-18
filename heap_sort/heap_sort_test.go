package heap_sort

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{0, 0}, []int{0, 0}},
		{[]int{9}, []int{9}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{-4}, []int{-4}},
		{[]int{3, -3, 4, -4, 1, -1, 7, -7}, []int{-7, -4, -3, -1, 1, 3, 4, 7}},
		{[]int{1, 2, 3, 1, 2, 3, 1, 2, 3}, []int{1, 1, 1, 2, 2, 2, 3, 3, 3}},
		{[]int{4, 5, 3}, []int{3, 4, 5}},
		{[]int{6, 1, 5, 9, 4, 2, 3, 7, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{6, 0, 1, 5, 9, 4, 2, 3, 7, 8}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	} {
		t.Run("", func(t *testing.T) {
			if HeapSort(tc.s); !reflect.DeepEqual(tc.s, tc.want) {
				t.Errorf("s = %v, got = %v", tc.s, tc.want)
			}
		})
	}

}
