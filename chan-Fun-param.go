package main

import "fmt"

func recivec(ping chan<- string,msg string) {

	ping <- msg
}

func sendrecchan(recv <- chan string, sendto chan <- string) {
	msg := <-recv
	sendto <- msg

}

func main() {

	ping:= make(chan string,1)
	pong := make(chan  string,1)

	recivec(ping,"Graffelo")
	//fmt.Println(<-ping)
	sendrecchan(pong,ping)
	fmt.Println(<-pong)

}

