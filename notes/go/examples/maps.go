package main

import (
	"fmt"
)

func main() {

	/* Declare a map */
	my_map := make(map[string]int)
	fmt.Println("Empty map:", my_map)

	/* Set values */
	my_map["this_year"] = 2014
	my_map["last_year"] = 2013
	fmt.Println("After setting values:", my_map)

	/* Delete value */
	delete(my_map, "this_year")
	fmt.Println("After deleting value:", my_map)

	/* Getting both return values*/
	val1, val2 := my_map["last_year"]
	fmt.Println("First return value of get:", val1)
	fmt.Println("Second return value of get:", val2)

	simple_map := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Simply declared map:", simple_map)

}
