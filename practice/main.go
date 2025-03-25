package main

import (
	"fmt"
)

const aConst int = 4

func main() {

	var text string = "A line of text"
	fmt.Println(text)
	fmt.Printf("Variable's type is %T\n", text) //Printf does not have a line feed

	var number int = 43
	fmt.Println(number)
	fmt.Printf("Variable's type is %T\n", number)

	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("Variable's type is %T\n", defaultInt)

	var anotherString string = "This is another string"
	fmt.Println(anotherString)
	fmt.Printf("Variable's type is %T\n", anotherString)

	colonString := "Colon string"
	fmt.Println(colonString)
	fmt.Printf("Variable's type is %T\n", colonString)

	fmt.Println(aConst)
	fmt.Printf("Variable's type is %T\n", aConst)

}
