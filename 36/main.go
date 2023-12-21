package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing work!")
		}
	}
}

func main() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 5)

	close(done)
}
