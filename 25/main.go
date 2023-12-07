package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	now := time.Now()
	date := time.Date(2019, time.December, 21, 05, 42, 0, 0, time.Local)

	fmt.Println(now.Unix())                                   // seconds
	fmt.Println(now.UnixMilli())                              // milliseconds
	fmt.Println(now.UTC().Format("2006-01-02T15:04:05.999Z")) // ISO String Date
	fmt.Println(now.Compare(date.Local()))                    // 1 -> future
	fmt.Println(date.Local().Compare(now))                    // -1 -> expired

	formatDuration, err := time.ParseDuration("1h29m0s")
	if err != nil {
		fmt.Println(err)
	}
	future := now.Add(formatDuration)

	fmt.Println(future)

	randInt := rand.Intn(2) // 0 <= randInt < 2
	fmt.Println(randInt)

	foods := []string{"Rice", "Garri", "Cassava", "Ewa", "Eba", "Beans"}
	choice := foods[rand.Intn(len(foods))]
	fmt.Println(choice)
}
