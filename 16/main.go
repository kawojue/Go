package main

import (
	"fmt"
	"strings"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   *float64
}

func newBill(name string, items map[string]float64, tip *float64) Bill {
	return Bill{
		items: items,
		name:  name,
		tip:   tip,
	}
}

func toTitle(s string) string {
	splitted := strings.Split(s, "")
	formatted := strings.ToUpper(splitted[0]) + strings.Join(splitted[1:], "")

	return formatted
}

func main() {
	bill1 := newBill("Raheem", map[string]float64{"egg": 5.99, "bread": 3.69}, nil)
	fmt.Printf("%s\n%v\n%v\n", bill1.name, bill1.items, bill1.tip)

	fmt.Println()

	var tipValue float64 = 0.33
	bill2 := newBill("Kawojue", map[string]float64{"yam": 7.99}, &tipValue)
	fmt.Printf("Name: %s\nItems: %v\nTip: %v\n", bill2.name, bill2.items, *bill2.tip)

	editBill1 := bill1.editName("muyiwa")
	fmt.Println(editBill1.name)

	fmt.Println()

	fmt.Println(editBill1.formatName())

	fmt.Println()

	bill1.billing()

	fmt.Println()

	bill2.billing()
}

func (bill Bill) editName(name string) Bill {
	b := Bill{
		name:  name,
		tip:   bill.tip,
		items: bill.items,
	}
	return b
}

func (bill Bill) formatName() Bill {
	b := Bill{
		name:  toTitle(bill.name),
		items: bill.items,
		tip:   bill.tip,
	}
	return b
}

func (bill Bill) billing() {
	var total float64
	breakdowns := "Bill Breakdown:\n"

	for item, price := range bill.items {
		breakdowns += fmt.Sprintf("\t%s: $%g\n", toTitle(item), price)
		total += price
	}

	fmt.Println(breakdowns)
	if bill.tip != nil {
		fmt.Printf("Total: $%v\n", total+*bill.tip)
	} else {
		fmt.Printf("Total: $%v\n", total)
	}
}
