package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	primeArr := []int{2, 3, 5, 7, 11}
	for _, prime := range primeArr {
		strPrime := strconv.Itoa(prime)
		_, err := file.WriteString(strPrime + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	file, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		fmt.Println("Prime: ", scan.Text())
	}
}
