package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	channel := make(chan struct{})

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("finished")
				channel <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("doing some work")
			}
		}
	}()

	<-channel
}
