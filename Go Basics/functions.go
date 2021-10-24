package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Good morning %v \n", name)
}

func sayBye(name string) {
	fmt.Printf("Goodbye %v \n", name)
}

func cycleNames(name []string, f func(string)) {
	for _, val := range name {
		f(val)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("Talha Yaseen")
	sayGreeting("Abhinav Singh")
	sayBye("Talha Yaseen")
	fmt.Println()
	cycleNames([]string{"Abhinav", "Lalu", "Rabri"}, sayGreeting)

	a1 := circleArea(2.3)
	a2 := circleArea((13))
	fmt.Printf("Area of circle, a1 = %0.2f and a2 = %0.3f \n", a1, a2)
}
