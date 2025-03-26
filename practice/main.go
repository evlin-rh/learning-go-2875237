package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1

	var result string

	switch dow {
	case 1:
		result = "It's 1"
		// fallthroughÍ
	case 2:
		result = "It's 2"
	default:
		result = "It's sth else"
	}
	fmt.Println(result)

}
