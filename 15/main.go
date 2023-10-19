package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func modifyPerson(p *Person) {
	p.Name = "Raheem"
	p.Age = 20
}

func main() {
	num1 := 20
	num2 := &num1

	*num2 = 30

	// Print memory addresses
	fmt.Println(num2)  // Memory address of num1 (value pointed to by num2)
	fmt.Println(&num2) // Memory address of num2 (the pointer variable)
	fmt.Println(&num1) // Memory address of num1

	fmt.Println()

	// Print values
	fmt.Println(num1)  // 30 (value of num1)
	fmt.Println(*num2) // 30 (value pointed to by num2)

	fmt.Println()

	my_woman := Person{Name: "Mary", Age: 21}
	modifyPerson(&my_woman)
	fmt.Println(my_woman) // { Raheem 20 }
}
