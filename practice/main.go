package main

import (
	"fmt"
	"io"
	"os"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	fmt.Println("Network requests")

	content := "Hello From Go"
	file, err := os.Create("./fromString1.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters \n", length)
	defer file.Close()
	defer readFile("./fromString.txt")
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println(string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
