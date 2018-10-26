package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	fmt.Println("Try to send to channel where nobody would listen to")
	ch <- 1

	fmt.Println("This will NEVER prints")
}
