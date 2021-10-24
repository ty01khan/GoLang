package main

import (
	"fmt"
)

func main() {
	age := 56

	// booleans
	fmt.Println(age <= 87)
	fmt.Println(age <= 45)
	fmt.Println(age == 56)
	fmt.Println(age != 56)

	// if-else statements
	if age < 30 {
		fmt.Printf("age is less than 30 \n\n")
	} else if age < 40 {
		fmt.Printf("age is less than 40 \n\n")
	} else {
		fmt.Printf("age is not less than 40 \n\n")
	}

	var scores = []int{79, 58, 99, 94, 45}
	for in, val := range scores {
		if in == 1 || in == 3 {
			fmt.Println("contining at position", in)
			continue
		}

		if in > 3 {
			fmt.Println("breaking the loop at position: ", in)
			break
		}
		fmt.Printf("the value at index %v is %v \n", in, val)
	}
}
