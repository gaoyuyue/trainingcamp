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
	r1 := convert(2 * a, b)
	r2 := convert(10 * a + 1, b)
	if r1 != nil {
		newResult = append(r1, a)
	} else if r2 != nil {
		newResult = append(r2, a)
	}
	return newResult
}