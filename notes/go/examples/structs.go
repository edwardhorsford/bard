package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	grades []int
}

func main() {

	cj := Person{name: "CJ", age: 28}
	fmt.Println("Person:", cj)

	fmt.Println("Person's name:", cj.name)
	fmt.Println("Person's age:", cj.age)

	cj.grades = append(cj.grades, 99, 89, 45)
	fmt.Println("Person's grades:", cj.grades)

}
