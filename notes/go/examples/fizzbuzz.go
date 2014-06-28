package main

import "fmt"

// Checks if num is divisable by 3, 5, or both and returns appropriate value
func fizzbuzz(num int) string {
	var (
		result      string = ""
		divizBy3    bool   = num%3 == 0
		divizBy5    bool   = num%5 == 0
		divizByBoth bool   = divizBy3 && divizBy5
	)

	if divizByBoth {
		result = "FizzBuzz"
	} else if divizBy3 {
		result = "Fizz"
	} else if divizBy5 {
		result = "Buzz"
	}

	return result

}

func main() {
	var start int = 1
	var end int

	// Ask user for input
	fmt.Println("Enter number: ")

	// Get input from user
	fmt.Scanf("%d", &end)

	for start <= end {
		var fzbz = fizzbuzz(start)
		if len(fzbz) > 0 {
			fmt.Println(fzbz)
		} else {
			fmt.Println(start)
		}
		start++
	}

}
