package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// Initialise counter
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// Uses the range of slice
	for i := range colors {
		fmt.Println(colors[i])
	}

	// For Each, _ used to ignore the index
	for _, color := range colors {
		fmt.Println(color)
	}

	// Similar to while loop
	value := 1
	for value < 10 {
		fmt.Println(value)
		value++
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
		if sum > 200 {
			goto theEnd
		}
	}

theEnd:
	fmt.Println("Broken")
}
