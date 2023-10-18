package main

import "fmt"

// Pass-by-value
func modifyValue(lastName string) {
	lastName = "Raheem"
}

func updateName(name string) string {
	name = "Muyiwa"
	return name
}

// Pass-by-reference (Passing a Pointer)
func modifyRef(name *string) {
	*name = "Raheem"
}

func main() {
	lastName := "Kawojue"
	modifyValue(lastName)
	fmt.Println(lastName) // Kawojue

	fmt.Println()

	middleName := "Kawojue"
	middleName = updateName(middleName)
	fmt.Println(middleName) // Muyiwa

	fmt.Println()

	firstName := "Kawojue"
	modifyRef(&firstName)
	fmt.Println(firstName) // Raheem
}
