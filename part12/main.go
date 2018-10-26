package main

import (
	"fmt"
	"sync"
)

var (
	x     int32
	wg    sync.WaitGroup
	mutex *sync.Mutex
)

func main() {
	fmt.Println("With atomic operation, we provide proper data synchronization")

	mutex = &sync.Mutex{}
	x = 0
	for i := 0; i < 10000; i++ {
		go Increment(&wg, mutex)
		go Decrement(&wg, mutex)
		wg.Add(2)
	}

	wg.Wait()
	fmt.Printf("[Final] x = %d\n", x)
}

func Increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	x = x + 1
	fmt.Printf("x = %d\n", x)
	mutex.Unlock()
	wg.Done()
}

func Decrement(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	x = x - 1
	fmt.Printf("x = %d\n", x)
	mutex.Unlock()
	wg.Done()
}
