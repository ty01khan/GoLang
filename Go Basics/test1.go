package main

import "fmt"

func main() {
	name := "Talha"
	age := 22

	fmt.Println("1. My name is", name, "and my age is", age)
	fmt.Println()

	/**
	 * Printf (formatted strings) %_ = format specifier
	 */
	fmt.Printf("2. My name is %v and my age is %v \n", name, age)
	fmt.Printf("3. Age is of type %T \n", age)
	fmt.Printf("4. you scored %f points! \n", 98.5)
	fmt.Printf("5. you scored %0.1f points! \n\n", 98.93)

	/**
	 * Sprintf (save formatted strings)
	 */
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is:", str)
}
