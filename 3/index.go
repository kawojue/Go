package main

import (
	"fmt"
)

func main() {
	happy()
	fmt.Println(dance("Kawojue"))
}

func dance(name string) string {
	return fmt.Sprintf("%s dances", name)
}

func happy() {
	fmt.Println("Once again, I love Golang!")
}
