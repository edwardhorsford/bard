package main

import (
	"fmt"
)

func main() {

	/* Declare slice */
	numbers := []int{1, 2, 3, 4, 5}

	/* Use range */
	for _, num := range numbers {
		fmt.Println(num)
	}

}
