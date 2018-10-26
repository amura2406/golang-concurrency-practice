package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		go PhantomSteal(&wg)
		wg.Add(1)
	}

	fmt.Println("Waiting for all tuyuls to finish")
	wg.Wait()
	fmt.Println("All tuyuls have finished !")
}

func PhantomSteal(wg *sync.WaitGroup) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("Steal successful, got %d golds\n", rand.Intn(1000))
	wg.Done()
}
