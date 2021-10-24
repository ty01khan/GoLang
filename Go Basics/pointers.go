package main

import "fmt"

func updateStringByPointer(x *string) {
	*x = "Yaseen"
}

func main() {
	name := "Talha"
	updateStringByPointer(&name)
	fmt.Println(name)
}
