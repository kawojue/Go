package main

import (
	stuff "example/project/mypkg"
	"fmt"
	"reflect"
)

func main() {
	fmt.Printf("Hello, %s!\n", stuff.Name)

	fmt.Println(stuff.IntArrToStrArr(1, 2, 3, 4, 5))

	titledName := stuff.ToTitle("raheem kawojue")
	fmt.Println(titledName)
	fmt.Println(reflect.TypeOf(titledName))
}
