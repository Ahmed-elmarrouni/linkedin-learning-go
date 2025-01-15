package main

import (
	"fmt"
)

func main() {

	anInt := 42
	var p = &anInt
	fmt.Println("value of p is: ", *p)

	value1 := 50
	pointer1 := &value1
	fmt.Println("value of pointer1 is: ", *pointer1)

	*pointer1 = *pointer1 / 15
	fmt.Println("value of pointer1 is: ", *pointer1)
	fmt.Println("value of value1 is: ", value1)

}
