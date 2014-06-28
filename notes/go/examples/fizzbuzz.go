package main

import "fmt"

// Checks if num is divisable by 3, 5, or both and returns appropriate value
func fizzbuzz(num int) {
	if num%3 == 0 && num%5 == 0 {
		return "FizzBuzz"
	} else if num%3 == 0 {
		return "Fizz"
	} else if num%5 == 0 {
		return "Buzz"
	} else {
		return num
	}
}

func main() {
	var start int = 0
	var end int

	// Ask user for input
	fmt.Println("Enter number: ")

	// Get input from user
	fmt.Scanf("%d", &end)

	for start <= end {
		fmt.Println(fizzbuzz(start))
		start = start + 1
	}

}
