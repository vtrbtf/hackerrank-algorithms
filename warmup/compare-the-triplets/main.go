package main

import (
	"fmt"
)

func scanToArray() []int {
	x := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scanf("%d", &x[i])
	}
	return x
}

func calculate(x1, x2 []int) int {
	r := 0
	for i := 0; i < 3; i++ {
		if x1[i] > x2[i] {
			r = r + 1
		}
	}
	return r
}

func main() {
	a := scanToArray()
	b := scanToArray()

	fmt.Print(calculate(a, b))
	fmt.Print(" ")
	fmt.Println(calculate(b, a))
}
