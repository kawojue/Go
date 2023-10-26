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

type Student struct {
	age    uint8
	grades map[string]float64
	name   string
}

func main() {
	c1 := Circle{3.4, &Point{y: 2, x: 3}}
	fmt.Println(c1.x)

	s1 := Student{age: 20, name: "Raheem", grades: map[string]float64{
		"English":     89.3,
		"Mathematics": 96.7,
		"Physics":     92,
		"Chemistry":   80.5,
	}}

	fmt.Printf("\tAgerage: %g%%\n", s1.getAverageGrade())
	fmt.Println(s1.getLargestGrade())
}

func (student Student) getAverageGrade() float64 {
	var avg float64 = 0

	for subject, grade := range student.grades {
		fmt.Printf("%s: %g%%\n", subject, grade)
		avg += grade
	}

	return avg / float64((len(student.grades)))
}

func (student Student) getLargestGrade() string {
	currentSubject := ""
	var currentMax float64 = 0

	for subject, grade := range student.grades {
		if grade > currentMax {
			currentMax = grade
			currentSubject = subject
		}
	}

	return fmt.Sprintf("The Subject with the Highest Grade is %s with %v%%\n", currentSubject, currentMax)
}
