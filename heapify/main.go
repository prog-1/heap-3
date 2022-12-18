package main

func heapify(s []int) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		heapifyUp(s, i)
	}
}

func heapifyUp(s []int, i int) {
	for {
		c := i*2 + 1
		if c+1 < len(s) && s[c+1] > s[c] {
			c = c + 1
		}
		if c >= len(s) || s[i] >= s[c] {
			return
		}
		s[i], s[c] = s[c], s[i]
		i = c
	}
}
