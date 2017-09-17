package main

import (
	"fmt"
	"sort"
)

func scanToArray() []int {
	x := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &x[i])
	}
	return x
}

func sumBetween(x []int) int {
	r := 0
	for index := 1; index < len(x)-1; index++ {
		r = r + x[index]
	}
	return r
}

func main() {
	x := scanToArray()
	sort.Ints(x)

	fmt.Print(sumBetween(x) + x[0])
	fmt.Print(" ")
	fmt.Println(sumBetween(x) + x[len(x)-1])
}
