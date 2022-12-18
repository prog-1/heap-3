package main

import (
	"math"
)

// func main() {
// 	s := []int{10, 1, 12, 100, 20, 2, 17}
// 	fmt.Println(MaxHeapify(s)) // 100 20 17 1 10 2 12
// }

func MaxHeapify(s []int) []int {
	// take last parent
	// swap it with greatest child
	// do it with each element from i/2-1 til 0
	// Heapify down the element that previously was the root

	if len(s) < 2 {
		return s
	}

	p := s[len(s)/2-1] // last parent
	c1 := s[len(s)-1]
	c2 := math.MinInt
	if len(s)%2 != 0 {
		c2 = s[len(s)-2]
	}
	if c1 > p && c1 > c2 {
		s[len(s)/2-1], s[len(s)-1] = s[len(s)-1], s[len(s)/2-1]
	} else if c2 > p {
		s[len(s)/2-1], s[len(s)-2] = s[len(s)-2], s[len(s)/2-1]
	}

	for i := len(s)/2 - 2; i > 0; i-- {
		p := s[i]
		l := s[i*2+1]
		r := s[i*2+2]
		if l > p && l > r {
			s[i], s[i*2+1] = s[i*2+1], s[i]
		} else if r > p {
			s[i], s[i*2+2] = s[i*2+2], s[i]
		}
	}

	for i := 0; i <= len(s)/2-1; {
		p := s[i]
		l := s[i*2+1]
		r := s[i*2+2]
		if l > p && l > r {
			s[i], s[i*2+1] = s[i*2+1], s[i]
			i = i*2 + 1
		} else if r > p {
			s[i], s[i*2+2] = s[i*2+2], s[i]
			i = i*2 + 2
		} else {
			break
		}
	}

	return s

}
