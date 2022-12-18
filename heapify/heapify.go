package heapify

func Parent(i int) int {
	return (i - 1) / 2
}
func L(i int) int {
	return 2*i + 1
}
func R(i int) int {
	return 2*i + 2
}

func heapifyUp(i int, s []int) {
	largest := i
	if L(i) < len(s) && s[L(i)] > s[largest] {
		largest = L(i)
	}
	if R(i) < len(s) && s[R(i)] > s[largest] {
		largest = R(i)
	}
	if largest != i {
		s[i], s[largest] = s[largest], s[i]
		heapifyUp(largest, s)
	}
}

func heapify(s []int) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		heapifyUp(i, s)
	}
}
