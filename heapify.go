package main

import "fmt"

func main() {
	s := []int{7, 9, 2, 15, 10, 5, 12}
	heapify(s)
	fmt.Println(s)
}

//turns slice into max heap in-place
func heapify(s []int) {

	lp := len(s)/2 - 1         //taking last parent
	for i := lp; i >= 0; i-- { // i = p
		max := i              //index of max from p,l,r
		if s[l(i)] > s[max] { // if l > p
			max = l(i)
		}
		if iv(s, r(i)) && s[r(i)] > s[max] { // if r > max (l or p)
			max = r(i)
		}
		if max != i { // if one or both c is > p
			s[i], s[max] = s[max], s[i] // swapping l or r with p
			heapifyUp(s, max)           //calling heapifyup for swapped child
		}
	}

}

//balances maxHeap from root to leaves (p > c)
func heapifyUp(s []int, i int) {
	for iv(s, l(i)) && s[i] < s[l(i)] || iv(s, r(i)) && s[i] < s[r(i)] { // while p < l or r
		max := i
		if iv(s, l(i)) && s[i] < s[l(i)] { // if l > p
			max = l(i)
		}
		if iv(s, r(i)) && s[i] < s[r(i)] { // if r > p
			max = r(i)
		}
		if max != i { // if l or r > p
			s[i], s[max] = s[max], s[i]
			//i = max
		}
	}
}

func l(i int) int            { return (i * 2) + 1 }
func r(i int) int            { return (i * 2) + 2 }
func iv(s []int, i int) bool { return i < len(s) } //is valid

/*
	HEAPIFY:

	Treat input slice as "broken" heap
	The point of heapify algorithm is to fix the "broken heap" in linear time.

	For each parent from last to first call heapifyUp.

	1. Take last parent and it's 2 childrens
	2. Take larger from childrens, and swap it with parent, if largest is > than parent
	3. Call heapifyup (from parent to childen) for swapped parent to children.

*/
