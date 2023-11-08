package main

import (
	"fmt"
	"time"
)

func main() {

	signalChannel := make(chan struct{})
	sendChannel := make(chan string)

	go func() {
		for {
			select {
			default:
				time.Sleep(time.Second)
				fmt.Println("doing some work...")
			case <-signalChannel:
				sendChannel <- "finished"
				return
			}
		}
	}()
	time.Sleep(time.Second * 3)

	close(signalChannel)
	fmt.Println(<-sendChannel)
}
