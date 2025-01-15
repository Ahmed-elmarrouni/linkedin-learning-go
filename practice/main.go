package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println("Structs")

	// states := make(map[string]string)
	// states["WA"] = "Washington"
	// states["OR"] = "Oregon"
	// states["TX"] = "Texas"
	// states["CA"] = "California"
	// states["TN"] = "Tennessee"
	// states["FL"] = "Florida"
	// states["GA"] = "Georgia"
	// states["SC"] = "South Carolina"

	// better way to declare a map
	states := map[string]string{
		"WA": "Washington",
		"OR": "Oregon",
		"TX": "Texas",
		"CA": "California",
		"TN": "Tennessee",
		"FL": "Florida",
		"GA": "Georgia",
		"SC": "South Carolina",
	}

	fmt.Println(states)

	fmt.Println(len(states))

	fmt.Println("----------------------------------------")

	delete(states, "TX")
	// fmt.Println(len(states))

	states["NY"] = "New York"
	fmt.Println(states)

	fmt.Println("----------------------------------------")

	for k, v := range states {
		fmt.Println(k, v)

	}

	fmt.Println("----------------------------------------")

	for k, v := range states {

		fmt.Printf("%v: %v\n", k, v)
	}

	fmt.Println("----------------------------------------")

	keys := make([]string, len(states))
	i := 0

	for k := range states {
		keys[i] = k
		i++
	}

	fmt.Println("Before sorting", keys)
	sort.Strings(keys)
	fmt.Println("After sorting", keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
