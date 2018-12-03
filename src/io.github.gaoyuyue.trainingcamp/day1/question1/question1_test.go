package main

import (
	"testing"
)

func TestCalMaxConLen(t *testing.T) {
	var tests = []struct {
		nums []int
		maxSize int
	}{
		{[]int{1, 3, 4, 5, 6, 9}, 2},
		{[]int{998, 999, 1000}, 2},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{33}, 0},

	}

	for _, test := range tests {
		maxSize := calMaxConLen(test.nums)
		if maxSize != test.maxSize {
			t.Errorf("Error: calMaxConLen input: %v expect: %v actual: %v ", test.nums, test.maxSize, maxSize)
		}
	}
}

func TestCheckRange(t *testing.T)  {
	tests := []struct {
		num int
		result bool
	}{
		{-1, false},
		{10001, false},
		{2, true},
	}

	for _, test := range tests{
		result := checkRange(test.num)
		if result != test.result{
			t.Errorf("Error: checkRange input: %v expect: %v actual: %v ", test.num, test.result, result)
		}
	}
}