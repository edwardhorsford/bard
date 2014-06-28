package main

import (
	"fmt"
)

func main() {

	/* Declare a new slice */
	c := make([]string, 3)
	fmt.Println("Empty slice:", c)

	/* Append a value to a slice */
	c = append(c, "meow", "cats", "kitty litter")
	fmt.Println("After append():", c)

	/* Retrieve value from slice */
	var first_value = c[0]
	fmt.Println("Retrieve c[0]:", first_value)

	/* Set value on slice */
	c[1] = "zoom"
	fmt.Println("After setting value:", c)

}
