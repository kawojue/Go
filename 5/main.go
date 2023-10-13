package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input your text: ")
	scanner.Scan()
	text := scanner.Text()

	println(text)

	fmt.Print("Year you were born: ")
	scanner.Scan()
	year, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	fmt.Printf("You are %d years old.\n", calcAge(year))
}

func println(text string) {
	fmt.Printf("%s\n", text)
}

func calcAge(year int64) int64 {
	return 2023 - year
}
