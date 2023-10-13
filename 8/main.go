package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("How old are you? ")
		scanner.Scan()

		age, _ := strconv.ParseInt(scanner.Text(), 10, 32)

		if age < 18 {
			fmt.Println("You're underage.")
		} else if age >= 32 {
			fmt.Println("Comrade, you should be married by now.")
		} else {
			fmt.Println("Have fun!")
		}

		fmt.Println()
	}
}
