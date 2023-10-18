package main

import "fmt"

// Pass-by-value
func modifyValue(surname string) {
	surname = "Raheem"
}

// Pass-by-reference (Passing a Pointer)
func modifyRef(name *string) {
	*name = "Raheem"
}

func main() {
	surname := "Kawojue"
	modifyValue(surname)

	fmt.Println(surname) // Kawojue

	fmt.Println()

	firstName := "Raheem"
	modifyRef(&firstName)

	fmt.Println(firstName)
}
