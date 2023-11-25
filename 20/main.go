package main

import "fmt"

func main() {
	chars := "abcde"

	ascii := map[string]rune{}

	for _, char := range chars {
		ascii[string(char)] = rune(char)
	}

	fmt.Println(ascii)
}
