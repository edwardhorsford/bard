package main

import (
	"fmt"
)

/* Declare a function */
func add_em(x int, y int) int {
	return x + y
}

/* Declare a variadic function */
func add_many(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {

	add_em_value := add_em(10, 15)
	add_many_value := add_many(1, 2, 3, 4, 5)

	fmt.Println("Standard addition function:", add_em_value)
	fmt.Println("Variadic addition function:", add_many_value)
}
