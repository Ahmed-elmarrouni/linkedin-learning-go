package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// os.Stdout.Sync()
	// input, _ := reader.ReadString('\n')
	// fmt.Println("You entered:", input)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter ur full Name: ")
	Full_Name, _ := reader.ReadString('\n')
	fmt.Println("Ur Full name is:", Full_Name)
}
