package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	a := make([]int, n)
	b := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &b[i])
	}

	a = convert(n, a)

	for i := 0; i < m; i++ {
		f := binarySearch(a, b[i])
		var k int
		if f > 0 {
			k = b[i] - a[f-1]
		} else {
			k = b[i]
		}
		fmt.Printf("%d %d \n", f+1, k)
	}
}

func convert(n int, a []int) []int {
	sum := 0
	for i := 0; i < n; i++ {
		t := a[i]
		a[i] += sum
		sum += t
	}
	return a
}

func binarySearch(arr []int, e int) int {
	begin := 0
	end := len(arr) - 1
	var middle int

	for begin <= end {
		middle = (end - begin) >> 1 + begin
		if e < arr[middle] {
			end = middle - 1
		} else if e > arr[middle] {
			begin = middle+1
		} else {
			return middle
		}
	}
	return begin
}