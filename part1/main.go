package main

import (
	"fmt"
	"time"
)

var ch chan bool

func main() {
	fmt.Println("Hello World")

	ch = make(chan bool)

	go StartAnotherRoutine()
	fmt.Println("Waiting for 5 secs")
	time.Sleep(5 * time.Second)
	<-ch
	fmt.Println("Message received")
	fmt.Println("Wait for 5 more secs")
	time.Sleep(5 * time.Second)
	fmt.Println("Finished !")
}

func StartAnotherRoutine() {
	fmt.Println("[Another] is going to send ...")
	ch <- true

	fmt.Println("[Another] has finished sending")
}
