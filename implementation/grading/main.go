package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scan(&n)
	r := make([]int, n)

	for i := 0; i < n; i++ {
		var current_grade int
		fmt.Scan(&current_grade)

		if current_grade >= 38 && (5-(current_grade%5)) < 3 {
			current_grade = current_grade + (5 - (current_grade % 5))
		}

		r[i] = current_grade
	}

	for _, grade := range r {
		fmt.Println(grade)
	}

}
