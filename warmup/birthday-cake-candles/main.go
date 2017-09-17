package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scan(&n)

	current_max := 0
	count := 0

	for i := 0; i < n; i++ {
		var current_n int
		fmt.Scanf("%d", &current_n)
		if current_n == current_max {
			count = count + 1
		}

		if current_n > current_max {
			current_max = current_n
			count = 1
		}

	}
	fmt.Println(count)
}
