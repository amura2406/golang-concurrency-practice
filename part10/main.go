package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	singletonInit := GetInstance()
	for i := 0; i < 1000; i++ {
		go func() {
			s := GetInstance()
			if singletonInit != s {
				fmt.Println("Not singleton !")
				os.Exit(-1)
			}
			wg.Done()
		}()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("Singleton check successful")
}

type singleton struct {
}

var (
	instance *singleton
	wg       sync.WaitGroup
	once     sync.Once
)

func GetInstance() *singleton {
	// if instance == nil {
	// 	instance = &singleton{}
	// }
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
