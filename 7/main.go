package main

import "fmt"

func main() {
	x := 0

	for x < 5 {
		fmt.Printf("Value of x: %d\n", x)
		x++
	}

	fmt.Println()

	scores := []int{1, 4, 6, 2, 3}

	for i := 0; i < len(scores); i++ {
		fmt.Printf("score [%d]: %d\n", i, scores[i])
	}

	fmt.Println()

	names := []string{"Muyiwa", "Raheem", "Mary"}

	for index, value := range names {
		fmt.Printf("The name at index %d is %s.\n", index, value)
	}
}
