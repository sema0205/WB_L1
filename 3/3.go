package main

import (
	"fmt"
	"sync"
)

func main() {

	massive := [5]int{2, 4, 6, 8, 10}

	var sum int

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 5; i++ {
		index := i
		wg.Add(1)
		go func(index int) {
			mu.Lock()
			sum += massive[index] * massive[index]
			mu.Unlock()
			wg.Done()
		}(index)
	}

	wg.Wait()
	fmt.Println(sum)
}
