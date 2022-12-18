package heapify

import (
	"reflect"
	"testing"
)

func isHeap(s []int, i int) bool {
	if i <= (len(s)-1)/2 {
		return true
	}
	if (s[i] > s[L(i)] && s[i] > s[R(i)]) && (isHeap(s, L(i)) && isHeap(s, R(i))) {
		return true
	}
	return false
}

func TestHeapify(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		// we dont really care about exact values here
		// only if result is max heap to sort array correctly
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{0, 0}, []int{0, 0}},
		{[]int{9}, []int{9}},
		{[]int{4, 5, 3}, []int{3, 4, 5}},
		{[]int{6, 1, 5, 9, 4, 2, 3, 7, 8}, []int{8}},
		{[]int{6, 0, 1, 5, 9, 4, 2, 3, 7, 8}, []int{9}},
	} {
		t.Run("", func(t *testing.T) {
			if Heapify(tc.s); !reflect.DeepEqual(tc.s, tc.want) {
				t.Errorf("s = %v, want = %v, but != max heap", tc.s, tc.want)
			}
		})
	}

}
