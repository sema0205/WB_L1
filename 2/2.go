package main

import (
	"fmt"
	"sync"
)

func main() {

	massive := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		index := i
		wg.Add(1)
		go func(index int) {
			fmt.Println(massive[index] * massive[index])
			wg.Done()
		}(index)
	}

	wg.Wait()
}
