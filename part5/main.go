package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int, 1)

	go StartAnotherRoutine(ch)
	go StartAnotherRoutine(ch)

	fmt.Println("Start listening buffered channel...")
	time.Sleep(3 * time.Second)
	fmt.Printf("Receive: %d\n", <-ch)
	time.Sleep(3 * time.Second)
	fmt.Printf("Receive: %d\n", <-ch)
	fmt.Println("Finished")
}

func StartAnotherRoutine(ch chan<- int) {
	ch <- rand.Intn(100)
	fmt.Println("Sending finished !")
}
