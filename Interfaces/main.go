package main

import (
	"fmt"
	"math"
)

// shape interfaces
type shape interface {
	circumfrence() float64
	area() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// square methods
func (s square) circumfrence() float64 {
	return 4 * s.length
}

func (s square) area() float64 {
	return s.length * s.length
}

// circle methods
func (c circle) circumfrence() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("Circumfrence of %T is: %0.2f \n", s, s.circumfrence())
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
}

func main() {
	shapes := []shape{
		square{length: 10},
		square{length: 5.2},

		circle{radius: 3.5},
		circle{radius: 10},
	}

	for _, val := range shapes {
		printShapeInfo(val)
		fmt.Printf("-----------\n\n")
	}
}
