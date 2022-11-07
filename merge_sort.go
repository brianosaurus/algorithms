package main

import (
	"fmt"
)

func merge(a []int, b []int) []int {
	c := make([]int, len(a)+len(b))
	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}
		k++
	}

	for i < len(a) {
		c[k] = a[i]
		i++
		k++
	}


	for j < len(b) {
		c[k] = b[j]
		i++
		j++
	}

	return c
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func main() {
	fmt.Println(mergeSort([]int{1, 3, 2, 5, 4}))
}