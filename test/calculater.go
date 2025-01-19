package main

import (
	"fmt"
	"strconv"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
	// Your code goes here.
	value1 := convertInputToValue(input1)
	value2 := convertInputToValue(input2)

	if operation == "+" {
		return addValues(value1, value2)
	} else if operation == "-" {
		return subtractValues(value1, value2)
	} else if operation == "*" {
		return multiplyValues(value1, value2)
	} else if operation == "/" {
		return divideValues(value1, value2)
	} else {
		fmt.Println("Wa si bouazza we don't have this operation")
		return 0
	}

}

func convertInputToValue(input string) float64 {

	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
