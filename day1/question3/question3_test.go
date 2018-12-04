package main

import (
	"github.com/gaoyuyue/trainingcamp/day1/common"
	"testing"
)


func TestConvert(t *testing.T) {
	tests := []struct{
		a int
		b int
		result []int
	}{
		{2, 162, []int{162, 81, 8, 4, 2}},
		{4, 42, nil},
		{100, 40021, []int{40021, 4002, 2001, 200, 100}},
	}
	
	for _,test := range tests {
		if result := convert(test.a, test.b); !common.Equals(result, test.result) {
			t.Errorf("Error: convert input: %v %v expect: %v actual: %v ", test.a, test.b, test.result, result)
		}
	}
}

func TestAddOne(t *testing.T) {
	tests := []struct {
		result    []int
		e         int
		newResult []int
	}{
		{[]int{1, 2, 3,}, 4, []int{1, 2, 3, 4,}},
		{[]int{}, 1, []int{1}},
	}

	for _, test := range tests {
		if newResult := addOne(test.result, test.e); !common.Equals(newResult, test.newResult) {
			t.Errorf("Error: addOne input: %v %v expect: %v actual: %v ", test.result, test.e, test.newResult, newResult)
		}
	}
}

