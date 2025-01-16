package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println(color)
	}

	value := 1
	for value < 100 {
		fmt.Println("value", value)
		value++
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("sum", sum)

		if sum > 200 {
			goto theEnd
		}
	}
theEnd:
	fmt.Println("The end")
}
