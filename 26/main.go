package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	// Encoding
	person := Person{Name: "Raheem", Age: 20, City: "Lagos"}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonData))

	// Decoding
	var newPerson Person
	if err = json.Unmarshal(jsonData, &newPerson); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Println("Decoded person:", newPerson)
}
