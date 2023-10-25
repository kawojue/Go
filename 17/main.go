package main

import "fmt"

type Point struct {
	x int64
	y int64
}

type Circle struct {
	radius float64
	*Point
}

func main() {
	c1 := Circle{3.4, &Point{y: 2, x: 3}}
	fmt.Println(c1.x)
}
