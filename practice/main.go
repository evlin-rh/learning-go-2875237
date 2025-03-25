package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv" //	String conversion types
	"strings" //	All packages for manipulating strings
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", input)

	fmt.Print("Enter number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Values of numbers: ", aFloat)
	}

}
