package main

import (
	"fmt"
	"sync"
)

func main() {

	massive := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var wg sync.WaitGroup

	values := make(chan int)
	changedValues := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(massive); i++ {
			values <- massive[i]
		}
		close(values)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for value := range values {
			changedValues <- value * 2
		}
		close(changedValues)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for value := range changedValues {
			fmt.Println(value)
		}
	}()

	wg.Wait()
}
