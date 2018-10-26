package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	dog1 := Bark("Woof!")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-dog1:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("Shut up daug !")
			return
		default:
			fmt.Println("I'm sick !!!")
		}
	}
}

func Bark(sound string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("[%d] %s", i, sound)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return ch
}
