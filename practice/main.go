package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	fmt.Print("enter a number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	// fmt.Println("You entered:", aFloat

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Number of value:", aFloat)
	}

}
