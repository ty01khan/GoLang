package main

import "fmt"

func main() {
	var ages = [3]int{56, 98, 73}
	fmt.Println(ages)
	fmt.Println()

	names := [4]string{"Talha", "Abhinav", "Avinash", "Rakesh"}
	fmt.Println(names[1])
	names[1] = "Yaseen"
	fmt.Println(names, len(names))
	fmt.Println()

	// Slices (use arrays under the bond)
	var scores = []int{79, 58, 99, 94, 45}
	fmt.Println(scores[2])
	scores[2] = 65
	scores = append(scores, 68)
	fmt.Println(scores, len(scores))
	fmt.Println()

	// Slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)
	fmt.Println()
}
