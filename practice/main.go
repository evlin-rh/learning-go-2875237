package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Orange", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5, 5) //initial size of 5, cap size of 5
	numbers[0] = 12
	numbers[1] = 42
	numbers[2] = 52
	numbers[3] = 82
	numbers[4] = 72

	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)

}
