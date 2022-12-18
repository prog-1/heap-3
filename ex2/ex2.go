package main

import "fmt"

func heapify(h []int) {
	for j := len(h)/2 - 1; j >= 0; j-- {
		if len(h)-1 == j*2+1 && h[j] < h[j*2+1] {
			healifyup(h, j)
		} else if h[j] < h[j*2+1] || h[j] < h[j*2+2] {
			healifyup(h, j)
		}
	}
}

func healifyup(h []int, j int) {
	for {
		if j*2+1 > len(h)-1 && j*2+2 > len(h)-1 {
			return
		}
		if j*2+2 > len(h)-1 && j*2+1 == len(h)-1 {
			if h[j] < h[j*2+1] {
				h[j], h[j*2+1] = h[j*2+1], h[j]
			}
			return
		}
		if h[j] < h[j*2+1] || h[j] < h[j*2+2] {
			if h[j*2+1] > h[j*2+2] {
				h[j], h[j*2+1] = h[j*2+1], h[j]
				j = j*2 + 1
				continue
			} else {
				h[j], h[j*2+2] = h[j*2+2], h[j]
				j = j*2 + 2
				continue
			}
		}
		return
	}

}

func heapSort(s []int) {
	if len(s) <= 1 {
		return
	}
	heapify(s)
	for i := len(s); i > 0; i-- {
		s[0], s[i-1] = s[i-1], s[0]
		healifyup(s[:i-1], 0)
	}
}

func main() {
	s := []int{3, 9, 8, 1, -1, 0}
	heapSort(s)
	fmt.Println(s)
}
