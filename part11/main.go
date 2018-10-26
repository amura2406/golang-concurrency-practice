package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	x  int32
	wg sync.WaitGroup
)

func main() {
	fmt.Println("With atomic operation, we provide proper data synchronization")
	x = 0
	for i := 0; i < 10000; i++ {
		go Increment(&wg)
		go Decrement(&wg)
		wg.Add(2)
	}

	wg.Wait()
	fmt.Printf("[Final] x = %d\n", x)
}

func Increment(wg *sync.WaitGroup) {
	atomic.AddInt32(&x, 1)
	fmt.Printf("x = %d\n", x)
	wg.Done()
}

func Decrement(wg *sync.WaitGroup) {
	atomic.AddInt32(&x, -1)
	fmt.Printf("x = %d\n", x)
	wg.Done()
}
