package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input your text: ")
	scanner.Scan()

	text := scanner.Text()

	print(text)
}

func print(text string) {
	fmt.Printf("%s", text)
}
