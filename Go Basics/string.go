package main

import "fmt"

func updateString(name string) {
	name = "Yaseen"
}

func updateMenu(test map[string]float64) {
	test["coffee"] = 4.99
	test["pie"] = 4.95
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays, structs
	name := "Talha"
	updateString(name)
	fmt.Println(name)

	// group B types -> slices, maps, functions
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)
}
