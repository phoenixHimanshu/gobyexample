package main

import "fmt"

func main() {

	messages := make(chan string,2)

	go func() { messages <- "ping" }()
	go func() { messages <- "pong" }()

	msg := <-messages
	msg1 := <-messages
	fmt.Println(msg)
	fmt.Println(msg1)
}
