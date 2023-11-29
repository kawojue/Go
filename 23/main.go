package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	_, err = os.Stat(fmt.Sprintf("%s\\%s", cwd, "data.txt"))
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist.")
	}

}
