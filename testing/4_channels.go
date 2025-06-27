package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")

	newChan := make(chan string)
	newChan <- "ping"

	message := <-newChan

	fmt.Println("message", message)
}
