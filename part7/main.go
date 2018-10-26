package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	dog1 := Bark("Woof!")
	dog2 := Bark("Auuuuuuu")

	combined := FanIn(dog1, dog2)
	for i := 0; i < 10; i++ {
		fmt.Println(<-combined)
	}
	fmt.Println("Dogs Fan In has finished")
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

func FanIn(ch1, ch2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case s := <-ch1:
				ch <- s
			case s := <-ch2:
				ch <- s
			}
		}
	}()
	return ch
}
