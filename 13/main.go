package main

import "fmt"

type Amount struct {
	qty   int
	price float64
}

func main() {
	menu := map[string]Amount{
		"noodles": {1.5, 40},
		"yam":     {5.99, 23},
		"bread":   {2.32, 15},
		"egg":     {1.15, 300},
	}

	var total float64

	for item, amount := range menu {
		fmt.Printf("%s: $%.2f * %d = $%.2f\n", item, amount.price, amount.qty, amount.price*float64(amount.qty))

		total += amount.price * float64(amount.qty)
	}

	fmt.Printf("Total: $%g\n", total)

}
