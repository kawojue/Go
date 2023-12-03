package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	file := fmt.Sprintf("%s\\%s", cwd, "data.txt")
	_, err = os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist.")
	} else {
		f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		if _, err := f.WriteString("Hello, World!\n"); err != nil {
			log.Fatal(err)
		}
	}
}
