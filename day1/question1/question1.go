package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scanln(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
		if !checkRange(nums[i]) {
			fmt.Printf("Error: The Input Number: %d Out Of Range.\n", nums[i])
			os.Exit(-1)
		}
	}
	fmt.Println(nums)
	fmt.Println(calMaxConLen(nums[:]))
}

func checkRange(num int) bool {
	if num < 1 || num > 1000{
		return false
	}
	return true
}

func calMaxConLen(nums []int) int {
	size, maxSize, maxBegin, maxEnd, begin := 1, 1, 0, 0, 0
	n := len(nums)
	if n <= 1 {
		return 0
	}
	for i := 1; i < n; i++ {
		if nums[i] - nums[i-1] == 1 {
			size++
		} else {
			if size > maxSize {
				maxSize = size
				maxBegin = begin
				maxEnd = i - 1
			}
			size = 1
			begin = i
		}
	}
	if size > maxSize {
		maxSize = size
		maxBegin = begin
		maxEnd = n - 1
	}
	if maxBegin == 0 && nums[0] == 1 || maxEnd == n - 1 && nums[maxEnd] == 1000{
		return maxSize - 1
	}
	return  maxSize - 2
}