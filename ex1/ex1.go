package main

import (
	"fmt"
)

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

func main() {
	h := []int{2, 8, 9, 12, 4, 7, 1, 5}
	heapify(h)
	fmt.Println(h)
}
