package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	result := convert(a, b)
	if result != nil{
		fmt.Println("Yes")
		n := len(result)
		fmt.Println(n)
		for i := n - 1; i >= 0; i-- {
			fmt.Printf("%d ", result[i])
		}
	} else {
		fmt.Println("No")
	}
}

func convert(a int, b int) []int {
	if a == b {
		return []int{a}
	} else if a > b {
		return nil
	}

	var newResult []int

	if result := convert(2 * a, b); result != nil {
		newResult = addOne(result, a)
	}

	if result := convert(10 * a + 1, b); result != nil {
		newResult = addOne(result, a)
	}
	return newResult
}

func addOne(result []int, a int) []int {
	n := len(result)
	newResult := make([]int, n+1)
	copy(newResult, result)
	newResult[n] = a
	return newResult
}
