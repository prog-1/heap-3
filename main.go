package main

func heapify(s []int) []int {
	if len(s) < 2 {
		return s
	}
	i := len(s)/2 - 1
	for ; i >= 0; i-- {
		heapifyUp(s, i)
	}
	i = len(s)/2 - 1
	for ; i >= 0; i-- {
		l := 2*i + 1
		r := 2*i + 2
		if l < len(s) && s[l] > s[i] || r < len(s) && s[r] > s[i] {
			return heapify(s)
		}
	}
	return s

}

func heapifyUp(s []int, i int) []int {
	l := 2*i + 1
	r := 2*i + 2
	if l < len(s) && r < len(s) && s[r] > s[l] && s[r] > s[i] {
		s[r], s[i] = s[i], s[r]
	} else if l < len(s) && r < len(s) && s[l] > s[r] && s[l] > s[i] {
		s[l], s[i] = s[i], s[l]
	} else if l > len(s) && r < len(s) && s[r] > s[i] {
		s[r], s[i] = s[i], s[r]
	} else if l < len(s) && r > len(s) && s[l] > s[i] {
		s[l], s[i] = s[i], s[l]
	}
	return s
}
func heapifyUp2(s []int, i int) {
	L := i
	l := 2*i + 1
	r := 2*i + 2
	if l < len(s) && s[l] > s[i] {
		L = l
	}
	if r < len(s) && s[r] > s[L] {
		L = r
	}
	if L != i {
		s[i], s[L] = s[L], s[i]
		heapifyUp2(s, L)
	}
}

func heapSort(s []int) {
	heapify(s)
	x := len(s)
	for x > 1 {
		s[0], s[x-1] = s[x-1], s[0]
		x -= 1
		heapifyUp2(s[:x], 0)
	}
}
