package main

import (
	"fmt"
)

type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func main() {
	fmt.Println(getSumGen(5.4, 3))
}
