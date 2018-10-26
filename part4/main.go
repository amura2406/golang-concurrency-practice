package main

import (
	"fmt"
	"time"
)

var x int

func main() {
	fmt.Println("The danger of concurrency without coordination")
	x = 0
	for i := 0; i < 1000; i++ {
		go Increment()
		go Decrement()
	}

	time.Sleep(5 * time.Second)
	fmt.Printf("[Final] x = %d\n", x)
}

func Increment() {
	x++
	fmt.Printf("x = %d\n", x)
}

func Decrement() {
	x--
	fmt.Printf("x = %d\n", x)
}
