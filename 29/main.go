package main

import (
	"fmt"
	"time"
)

func printTo15() {
	for i := 1; i <= 15; i++ {
		fmt.Println("Func 1 : ", i)
	}
}

func printTo10() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Func 2 : ", i)
	}
}

func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}

func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

func main() {
	go printTo10()
	go printTo15()

	time.Sleep(2 * time.Second)

	channel1 := make(chan int)
	channel2 := make(chan int)

	go nums1(channel1)
	go nums2(channel2)

	fmt.Println(<-channel1)
	fmt.Println(<-channel1)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)
}
