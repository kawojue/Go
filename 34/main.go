package main

import "fmt"

func main() {
	myChannel := make(chan string)         // unbuffered channel
	anotherChannel := make(chan string, 0) // still unbuffered channel

	go func() {
		myChannel <- "data"
	}()

	go func() {
		anotherChannel <- "mydata"
	}()

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}
}
