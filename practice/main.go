package main

import (
	"fmt"
)

func main() {

	theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "Less 0"
	} else if theAnswer == 0 {
		result = "0"
	} else {
		result = "> 0"
	}
	fmt.Println(result)
}
