package main

import "fmt"

func main() {
	intSum := func(x int, y int) int { return x + y }

	fmt.Println(intSum(2, 3))
}
