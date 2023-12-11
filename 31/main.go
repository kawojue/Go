package main

import (
	"fmt"
)

func main() {
	intSum := func(x int, y int) int { return x + y }

	fmt.Println(intSum(2, 3)) // 5

	samp1 := 1
	func() {
		samp1 += 1
	}()

	fmt.Println("Samp 1:", samp1) // 2
	// this is not supposed to happen but cuz it is closure.
}
