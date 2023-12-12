package main

import (
	"fmt"
	"regexp"
)

func main() {
	reStr := "The ape was the the apex"
	match, err := regexp.MatchString("(ape [^ ]?)", reStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(reStr)
	fmt.Println(match)

	reStr2 := "cat rat fat mat pat"
	re, err := regexp.Compile("([pmrfc]at)")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reStr2)
	fmt.Println(re.MatchString(reStr2))       // returns boolean
	fmt.Println(re.FindString(reStr2))        // returns first match
	fmt.Println(re.FindAllString(reStr2, -1)) // returns all matches - slice
}
