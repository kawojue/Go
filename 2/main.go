package main

import (
	"fmt"
	"strings"
)

func main() {
	greetings := "Welcome, Are you there?"

	fmt.Println(strings.Count(greetings, "e"))
	fmt.Println(strings.ToUpper("I love Go!"))
	fmt.Println(strings.Contains(greetings, "A"))
	fmt.Println(strings.ToLower(fmt.Sprintf("%s", "Kawojue Raheem")))

	replacedSring := strings.ReplaceAll(greetings, "Welcome", "Hello")

	fmt.Println(strings.Index(replacedSring, "you"))
	fmt.Println(strings.Split(replacedSring, " "))

	fmt.Println(replacedSring)
}
