package main

import (
	"fmt"
)

func main() {

	anInt := 42
	var p = &anInt //* = shows its pointer not value
	fmt.Println("Value of p: ", *p)

	value1 := 41.23
	pointer1 := &value1
	fmt.Println("Value1: ", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Pointer1: ", *pointer1)
	fmt.Println("Value1: ", value1)

}
