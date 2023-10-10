package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	dance := dance("Kawojue")

	wg.Add(1)

	go func() {
		defer wg.Done()
		happy()
	}()

	fmt.Println(dance)

	wg.Wait()
}

func dance(name string) string {
	return fmt.Sprintf("%s dances", name)
}

func happy() {
	fmt.Println("Once again, I love Golang!")
}
