package main

/* Import packages */
import (
	"fmt"
	"os"
	"strings"
)

/* Declare a variable */
var who string = "World"

/* Declare a constant */
const name = "Meerkat"

/* Declare an array */
var array = []int{1, 2, 3, 4}

/* Define a slice */
var slice = array[0:2]

/* Declare a list of constants */
const (
	car      = "BMW"
	cat_name = "Mr.Meowingtons"
	city     = "The Big Apple"
)

/* Declare a function */
func add_em(x int, y int) int {
	return x + y
}

func main() {
	/* Declare a variable within a function */
	variable_in_function := "I'm in a function!"

	/* Declare a map */
	var my_map = make(map[string]string)
	my_map["cat"] = "Mr.Meow"

	/* Conditional */
	if len(os.Args) > 1 {
		/* os.Args[0] is "hello" or "hello.exe" */
		who = strings.Join(os.Args[1:], " ")
	}

	fmt.Println("Array:", array)
	fmt.Println("Slice:", slice)
	fmt.Println("Map:", my_map)
	fmt.Println("Map.cat:", my_map["cat"])
	fmt.Println("Varible In Function:", variable_in_function)
	fmt.Println("Function Return Value:", add_em(10, 15))
	fmt.Println("Name from command line:", "Hello", who+"!")
}
