package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	now := time.Now()
	date := time.Date(2019, time.December, 21, 05, 42, 0, 0, time.Local)

	fmt.Println(now.Unix())                              // seconds
	fmt.Println(now.UnixMilli())                         // milliseconds
	fmt.Println(time.Now().Local().Format("2006-01-02")) // ISOString Date
	fmt.Println(time.Now().Compare(date.Local()))        // 1 -> future
	fmt.Println(date.Local().Compare(time.Now()))        // -1 -> expired

	randInt := rand.Intn(2) // 0 <= randInt < 2
	fmt.Println(randInt)

	foods := []string{"Rice", "Garri", "Cassava", "Ewa", "Eba", "Beans"}
	choice := foods[rand.Intn(len(foods))]
	fmt.Println(choice)
}
