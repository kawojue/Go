package main

import (
	"fmt"
)

type MyConstraint interface {
	int | float64
}

type Customer struct {
	name string
	city string
	bal  float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func editCity(c *Customer, city string) {
	c.city = city
}

func main() {
	fmt.Println(getSumGen(5.4, 3))

	var newCustomer Customer
	newCustomer.name = "Raheem"
	newCustomer.city = "Lagos"
	newCustomer.bal = 6.99

	fmt.Println(newCustomer)

	editCity(&newCustomer, "Ogun")
	fmt.Println(newCustomer)
}
