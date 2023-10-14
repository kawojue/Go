package main

import (
	"fmt"
	"math"
	"strconv"
)

func sayGoodbye(name string) string {
	return fmt.Sprintf("Goodbye, %s.", name)
}

func cycleNames(names []string, cb func(name string) string) {
	for _, value := range names {
		fmt.Println(cb(value))
	}
}

func circleArea(radius float64) float64 {
	area := math.Pi * math.Pow(radius, 2)
	formatted, _ := strconv.ParseFloat(
		fmt.Sprintf("%.3f", area), 64,
	)
	return formatted
}

func main() {
	cycleNames([]string{"Raheem", "Muyiwa"}, sayGoodbye)

	fmt.Println(circleArea(12))
}
