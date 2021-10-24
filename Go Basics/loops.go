package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("value of x is:")
	for x < 5 {
		fmt.Printf("\t\t%v \n", x)
		x++
	}

	fmt.Println("value of i is:")
	for i := 0; i < 5; i++ {
		fmt.Printf("\t\t%v \n", i)
	}
	fmt.Println()

	var scores = []int{79, 58, 99, 94, 45}
	sum := 0
	fmt.Printf("Sum of elements of slice is: ")
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	fmt.Printf("%v \n\n", sum)

	for i, v := range scores {
		fmt.Printf("the value at index %v is %v \n", i, v)
	}
}
