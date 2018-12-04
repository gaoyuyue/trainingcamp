package main

import (
	"github.com/gaoyuyue/trainingcamp/day1/common"
	"testing"
)

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

func TestConvert(t *testing.T) {
	tests := []struct {
		n      int
		a      []int
		result []int
	}{
		{3, []int{10, 15, 12}, []int{10, 25, 37}},
	}

	for _, test := range tests {
		if result := convert(test.n, test.a); !common.Equals(result, test.result) {
			t.Errorf("Error: convert input: %v %v except: %v actual: %v .", test.n, test.a, test.result, result)
		}
	}
}