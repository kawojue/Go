package main

import "fmt"

func main() {
	menu := map[string]float64{
		"egg":     1.15,
		"noodles": 1.5,
		"yam":     5.99,
	}

	menu["milo"] = 2.0
	menu["yam"] = 10.99

	fmt.Println(menu)
	fmt.Println(menu["noodles"])

	for item, price := range menu {
		fmt.Printf("%s: $%.2f\n", item, price)
	}

	delete(menu, "egg")

	amount, exist := menu["egg"]

	if exist {
		fmt.Printf("Egg is on the menu for $%.2f\n", amount)
	} else {
		fmt.Println("Egg is not on the menu.")
	}
}
