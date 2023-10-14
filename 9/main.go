package main

import (
	"fmt"
)

func sayGoodbye(name string) string {
	return fmt.Sprintf("Goodbye, %s.", name)
}

func callback(names []string, cb func(name string) string) {
	for _, value := range names {
		fmt.Println(cb(value))
	}
}

func main() {
	callback([]string{"Raheem", "Muyiwa"}, sayGoodbye)
}
