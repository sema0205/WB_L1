package main

import (
	"fmt"
	"sync"
)

type Iterator struct {
	mu    sync.Mutex
	value int
}

func (i *Iterator) Iterate(value int) {
	i.mu.Lock()
	i.value += value
	i.mu.Unlock()
}

func main() {

	var temp Iterator
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		index := i
		wg.Add(1)
		go func(val int) {
			for i := 0; i < 10; i++ {
				temp.Iterate(val)
			}
			wg.Done()
		}(index)
	}

	wg.Wait()
	fmt.Println(temp.value)
}
