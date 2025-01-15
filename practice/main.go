package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")

	var carsBrand [6]string
	carsBrand[0] = "Toyota"
	carsBrand[1] = "Ford"
	carsBrand[2] = "Chevrolet"
	carsBrand[3] = "Nissan"
	carsBrand[4] = "BMW"
	carsBrand[5] = "Mercedes Benz w123 ❤️"

	fmt.Println(carsBrand)
	fmt.Println(carsBrand[3])
	fmt.Println("Numbers of cars", len(carsBrand))

	// if i have unlimited number of elements "we called slices"
	carsModel := []string{"Corolla", "Focus", "Camaro"}
	fmt.Println(carsModel)

	// add new item to the slice
	carsModel = append(carsModel, "Dacia")
	fmt.Println(carsModel)

	// create a slice
	var colors = []string{"Red", "Green", "Blue"}
	colors = append(colors, "Yellow", "Purple")
	fmt.Println(colors)

	// remove the first item from the slice
	colors = append(colors[1:len(colors)])
	fmt.Println("remove the first an item", colors)

	// remove the last item from the slice
	colors = append(colors[:len(colors)-1])
	fmt.Println("remove the last item", colors)

	// remove the item from the middle of the slice
	colors = append(colors[:2], colors[3:]...)
	fmt.Println(colors)

}
