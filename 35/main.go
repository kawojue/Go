package main

import "fmt"

func main() {
	charChannel := make(chan string, 3) // buffered channel
	chars := []string{"a", "b", "c"}

	for _, char := range chars {
		select {
		case charChannel <- char:
		}

		// charChannel <- char
	}

	close(charChannel)

	for res := range charChannel {
		fmt.Println(res)
	}
}
