package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.49,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for key, val := range menu {
		fmt.Println(key, "-", val)
	}
	fmt.Println()

	// ints as key type
	phonebook := map[int]string{
		268489505: "Talha",
		759589696: "Abhinav",
		278489599: "Avinash",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[759589696])

	phonebook[759589696] = "Rajan"
	fmt.Println(phonebook[759589696])
}
