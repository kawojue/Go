package main

import "fmt"

func main() {
	fetch(1, 2, 3, 4, 5)
}

func fetch(sizes ...int) {
	for _, size := range sizes {
		fmt.Println(size)
	}
}
