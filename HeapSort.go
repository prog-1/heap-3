package main

func HeapSort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	s = MaxHeapify(s)
	// Swap root with the last element
	// Heapify down the new root

	for i := len(s) - 1; i > 0; i-- {
		s[i], s[0] = s[0], s[i]
		HeapifyDown(s[:i], 0)
	}
	return s
}

func HeapifyDown(s []int, in int) {
	// Until we're not at the bottom layer:
	// 	Take left and right children
	// 	If any of children is bigger than parent, then swap it with parent; else - break
	// 	Proceed with new parent's place
	if len(s) < 2 {
		return
	}
	for i := in; i < len(s)/2-1; {
		l := s[2*i+1]
		r := s[2*i+2]
		if l > s[i] && l > r {
			s[i], s[2*i+1] = s[2*i+1], s[i]
			i = 2*i + 1
		} else if r > s[i] {
			s[i], s[2*i+2] = s[2*i+2], s[i]
			i = 2*i + 2
		} else {
			return
		}
	}
	if len(s)%2 == 0 {
		if s[len(s)-1] > s[len(s)/2-1] {
			s[len(s)-1], s[len(s)/2-1] = s[len(s)/2-1], s[len(s)-1]
		}
	} else {
		i := len(s)/2 - 1
		l := s[i*2+1]
		r := s[i*2+2]
		if l > s[i] && l > r {
			s[i], s[2*i+1] = s[2*i+1], s[i]
		} else if r > s[i] {
			s[i], s[2*i+2] = s[2*i+2], s[i]
		}
	}
}
