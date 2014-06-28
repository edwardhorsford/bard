package main

import (
	"fmt"
	"os"
	"strings"
)

/*  Declare a constant */
const name = "Meerkat"

/* Declare a list of constants */
const (
	car      = "BMW"
	cat_name = "Mr.Meowingtons"
	city     = "The Big Apple"
)

func add_em(x int, y int) int {
	return x + y
}

func main() {
	who := "World!"

	/* Declare an array */
	var array = []int{1, 2, 3, 4}

	/* Ternary Operator */
	num := "cats"

	if len(os.Args) > 1 {
		/* os.Args[0] is "hello" or "hello.exe" */
		who = strings.Join(os.Args[1:], " ")
	}

	fmt.Println(num)
	fmt.Println(array[1:])
	fmt.Println(add_em(10, 15))
	fmt.Println("Hello", who+"!")
}
