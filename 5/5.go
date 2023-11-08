package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	secondsToStop := 5

	testChannel := make(chan string)
	signalChannel := make(chan struct{})
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		for {
			select {
			case <-signalChannel:
				wg.Done()
				return
			default:
				time.Sleep(time.Second * 1)
				message := <-testChannel
				fmt.Println("read", message)
			}
		}
	}()

	go func() {
		wg.Add(1)
		var counter int
		for {
			select {
			case <-signalChannel:
				wg.Done()
				return
			default:
				time.Sleep(time.Second * 1)
				counter++
				message := fmt.Sprintf("message %d", counter)
				testChannel <- message
				fmt.Println("posted", message)
			}
		}
	}()

	time.Sleep(time.Duration(secondsToStop) * time.Second)
	close(signalChannel)

	wg.Wait()
}
