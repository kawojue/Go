package main

import "fmt"

var name = "Kawojue" // outside main scope

func main() {
	age := 100

	var num uint8 = 255

	fmt.Println(num, age)
	fmt.Print(num, "\n")
	fmt.Println(name)
}
