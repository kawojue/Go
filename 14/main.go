package main

import "fmt"

func updateName(name string) {
	name = "Raheem"

	fmt.Println(name)
}

func main() {
	name := "Kawojue"

	updateName(name)

	fmt.Println(name)
}
