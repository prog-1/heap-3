package main

import "fmt"

func heapifyUP(s []int, i int) {
	for i < len(s) {
		left := i*2 + 1
		right := left + 1
		if left >= len(s) {
			return
		}
		if right < len(s) {
			if s[left] > s[right] {
				if s[left] > s[i] {
					s[left], s[i] = s[i], s[left]
					i = left
				} else {
					return
				}
			} else {
				if s[right] > s[i] {
					s[right], s[i] = s[i], s[right]
					i = right
				} else {
					return
				}
			}
		} else {
			if s[left] > s[i] {
				s[left], s[i] = s[i], s[left]
				i = left
			}
			return
		}
	}
}

func heapify(s []int) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		heapifyUP(s, i)
	}
}

func heapSort(s []int) {
	for len(s) > 1 {
		s[0], s[len(s)-1] = s[len(s)-1], s[0]
		s = s[:len(s)-1]
		heapify(s)
	}
}

func main() {
	a := []int{10, 9, 7, 8, 5, 6, 3, 1, 4, 2}
	heapSort(a)
	fmt.Println(a)
}
