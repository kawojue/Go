package main

import (
	"fmt"
	"strings"
)

func getInitals(name string) (f string, l string) {
	names := strings.Split(name, " ")

	formattedNames := []string{}

	for _, value := range names {
		splitV := strings.Split(value, "")
		formattedNames = append(
			formattedNames, (strings.ToUpper(splitV[0]) + strings.ToLower(strings.Join(splitV[1:], ""))),
		)
	}

	if len(formattedNames) > 1 {
		return formattedNames[0], formattedNames[1]
	}

	return formattedNames[0], "_"
}

func main() {
	firstName, lastName := getInitals("raheem kawojue")

	fmt.Println(firstName, lastName)
}
