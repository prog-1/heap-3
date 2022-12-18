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

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	heapify(a)
	fmt.Println(a)
}
