package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)

	var result string
	switch dow {
	case 1:
		result = "It's Sunday!"
	case 2:
		result = "It's Monday!"
	case 3:
		result = "It's Tuesday!"
	default:
		result = "It's a weekday!"
	}

	fmt.Println("Day", result)

}
