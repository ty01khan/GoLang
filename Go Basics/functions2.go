package main

import (
	"fmt"
	"strings"
)

func getInitials(name string) (string, string) {
	s := strings.ToUpper(name)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], ""
}

func main() {
	fn1, sn1 := getInitials("Talha Yaseen")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("Abhinav Kumar Singh")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("Irfan")
	fmt.Println(fn3, sn3)
}
