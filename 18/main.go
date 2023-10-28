package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rect struct {
	width  float64
	height float64
}

func main() {
	c1 := Circle{23}
	r1 := Rect{4, 8}

	shapes := []Shape{c1, r1}

	for _, shape := range shapes {
		printArea(shape)
	}
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func printArea(shape Shape) {
	fmt.Println(shape.area())
}
