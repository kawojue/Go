package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := input("What? ")
	fmt.Println(text)
}

func input(question string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	if question != "" {
		fmt.Print(question)
	}

	text, err := reader.ReadString('\n')

	return strings.TrimSpace(text), err
}
