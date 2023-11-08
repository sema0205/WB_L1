package main

import "time"

func sleep(duration time.Duration) {
	<-time.Tick(duration)
}

func main() {
	sleep(time.Second * 2)
}
