package main

import (
	"errors"
	"fmt"
	"time"
)

type Date struct {
	day   int
	month int
	year  int
}

func main() {
	newDate := Date{07, 12, 2023}
	fmt.Println(newDate)

	if err := newDate.SetMonth(11); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newDate)
	}

	if err := newDate.SetYear(13); err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(newDate)
	}

}

func (d *Date) SetDay(day int) error {
	if (day < 1) || (day > 31) {
		return errors.New("Invalid Day Provided.")
	}

	d.day = day
	return nil
}

func (d *Date) SetMonth(month int) error {
	if (month < 1) || (month > 12) {
		return errors.New("Invalid Month Provided.")
	}

	d.month = month
	return nil
}

func (d *Date) SetYear(year int) error {
	if (year < 1900) || (year > time.Now().Year()) {
		return errors.New("Invalid Year Provided.")
	}

	d.year = year
	return nil
}
