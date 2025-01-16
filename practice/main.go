package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fileConetent := "Hello, My name is Ahmed"
	file, err := os.Create("./ahmedFile.txt")
	checkErr(err)
	length, err := io.WriteString(file, fileConetent)
	checkErr(err)
	fmt.Printf("wrote file with %v characters\n", length)
	defer file.Close()

	defer readFile("./ahmedFile.txt")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkErr(err)
	fmt.Println("File content: ", string(data))
}
