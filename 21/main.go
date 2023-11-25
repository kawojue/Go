package main

import "fmt"

func main() {
	ans, err := getQuotient(6, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}

	getSum(1, 2, 3, 4, 5, 6, 7, 8, 9)
}

func getQuotient(x float64, y float64) (answer float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("Can't divide by zero!")
	}

	return x / y, nil
}

func getSum(nums ...int) {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)
}
