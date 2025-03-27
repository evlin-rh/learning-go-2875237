package main

import (
	"fmt"
)

func main() {
	sum := addValues(10, 20)
	fmt.Println("The sum is: ", sum)
	sum1 := addValue2(10, 10)
	fmt.Println("The sum is: ", sum1)
	doSomething()
	multiSum, multiCount := addAllValue(1, 2, 3, 4, 5)
	fmt.Println("Multivalue: ", multiSum, "MultiCount: ", multiCount)
}

func doSomething() {
	fmt.Println("Do sth")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

func addValue2(value1, value2 int) int {
	return value1 + value2
}

func addAllValue(values ...int) (int, int) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, len(values)
}
