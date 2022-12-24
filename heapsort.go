package main

import "fmt"

func main() {
	s := []int{4, 6, 2, 7, 1, 5, 3}
	heapsort(s)
	fmt.Println(s)
}

//sorts slice in-place in ascending order
func heapsort(s []int) {
	heapify(s)
	for i := len(s) - 1; i > 0; i-- {
		s[0], s[i] = s[i], s[0] //swap 1st with last
		s = s[:i]               //exclude last sorted element
		heapifyUp(s, 0)         //balance broken heap
	}
}
