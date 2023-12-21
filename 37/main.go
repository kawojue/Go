package main

import (
	"fmt"
)

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(nums <-chan int) <-chan int {
	out := make(chan int, 0)
	go func() {
		for n := range nums {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func main() {
	nums := []int{2, 5, 4, 7, 3, 1, 6}
	dataChannel := sliceToChannel(nums)

	finalChannel := sq(dataChannel)
	for n := range finalChannel {
		fmt.Println(n)
	}
}
