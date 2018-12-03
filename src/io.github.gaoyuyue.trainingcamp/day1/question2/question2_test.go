package main

import "testing"

func TestBinarySearch(t *testing.T)  {
	tests := []struct {
		arr    []int
		e      int
		result int
	}{
		{[]int{10, 25, 37}, 1, 0},
		{[]int{10, 25, 37}, 9, 0},
		{[]int{10, 25, 37}, 12, 1},
		{[]int{10, 25, 37}, 23, 1},
		{[]int{10, 25, 37}, 26, 2},
		{[]int{10, 25, 37}, 37, 2},
	}

	for _, test := range tests {
		if result := binarySearch(test.arr, test.e); result != test.result{
			t.Errorf("Error: binarySearch input: %v %v expect: %v actual: %v ", test.arr, test.e, test.result, result)
		}
	}
}
