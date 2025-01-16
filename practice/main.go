package main

import (
	"fmt"
)

func main() {
	doSomething()

	multiSum, multiCount := addAllValues(7, 8, 10)
	fmt.Println("Sum of all values is: ", multiSum)
	fmt.Println("Count of all values is: ", multiCount)
}

func doSomething() {
	fmt.Println("Doing something")
	sumV := addValues(8, 2)
	fmt.Println("Sum of Va and V2 is: ", sumV)
}

func addValues(valueA int, valueB int) int {

	return valueA + valueB
}

func addAllValues(value ...int) (int, int) {
	total := 0
	for _, val := range value {
		total += val
	}

	return total, len(value)
}
