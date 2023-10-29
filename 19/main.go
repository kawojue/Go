package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	text, _ := getInput("What? ", bufio.NewReader(os.Stdin))
	fmt.Println(text)

	promptOptions()
}

func promptOptions() {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("How old are you?", reader)
	age, _ := strconv.ParseInt(option, 10, 64)

	switch {
	case age >= 18:
		fmt.Println("You're Adult.")
	case age < 18:
		fmt.Println("You're underage.")
	default:
		fmt.Println("Invalid Age Provided.")
	}
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}
