package main

import (
	"fmt"
)

func main() {
	dance := dance("Kawojue")

	happy()
	fmt.Print(dance)
}

func dance(name string) string {
	return fmt.Sprintf("%s dances", name)
}

func happy() {
	fmt.Println("Once again, I love Golang!")
}
