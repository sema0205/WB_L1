package main

import (
	"fmt"
	"sync"
)

type MapMutex struct {
	wg        sync.Mutex
	container map[string]int
}

func (m *MapMutex) AddValue(key string, value int) {
	m.wg.Lock()
	m.container[key] = value
	m.wg.Unlock()
}

func main() {

	exampleContainer := MapMutex{
		wg:        sync.Mutex{},
		container: make(map[string]int),
	}

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(key string, value int) {
			exampleContainer.AddValue(key, value)
			wg.Done()
		}(fmt.Sprintf("%d", i), i)
	}

	wg.Wait()
	fmt.Println(exampleContainer.container)
}
