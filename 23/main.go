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

	filePath := fmt.Sprintf("%s\\%s", cwd, "data.txt")
	_, err = os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist.")
	} else {
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		if _, err := file.WriteString("Hello, World!\n"); err != nil {
			log.Fatal(err)
		}
	}
}
