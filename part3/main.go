package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go RunForever()
	fmt.Println("Try to send to channel where nobody would listen to")
	ch <- 1 // will not panic since there's another goroutine running, until no goroutine beside main, it'll panic

	fmt.Println("This will NEVER prints")
}

func RunForever() {
	for i := 0; i < 5; i++ {
		fmt.Println("Loop")
		time.Sleep(5 * time.Second)
	}
}
