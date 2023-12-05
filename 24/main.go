package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	iArgs := []int{}

	for _, i := range args {
		value, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, value)
	}

	fmt.Println(iArgs)

	max := 0
	for i := 0; i < len(iArgs); i++ {
		if iArgs[i] > max {
			max = iArgs[i]
		}
	}

	fmt.Printf("Max value: %v\n", max)
}

// go build main.go
// .\ main args
