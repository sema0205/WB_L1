package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func worker(index int, jobs <-chan string, wg *sync.WaitGroup) {
	for job := range jobs {
		fmt.Println("worker", index, "doing job", job)
	}
	wg.Done()
}

func main() {

	var workersAmount int
	fmt.Scanln(&workersAmount)

	var wg sync.WaitGroup

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	jobs := make(chan string, workersAmount)
	for i := 0; i < workersAmount; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	var counter int

	go func() {
		for {
			counter++
			testString := fmt.Sprintf("%d", counter)
			jobs <- testString
		}
	}()

	<-c

	wg.Wait()
	close(jobs)
}
