package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	go func() { messages <- "pong" }()

	go func() { messages <- "pow" }()

	msg := <-messages
	msg2 := <-messages
	msg3 := <-messages

	fmt.Println(msg)
	fmt.Println(msg2)
	fmt.Println(msg3)

}
