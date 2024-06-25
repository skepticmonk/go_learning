package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

// (c * Circle) is called as receiver,
// which facilitates the function to be used as a method of Circle
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func struct_example() {
	var c Circle
	c = Circle{x: 0, y: 0, r: 5}
	fmt.Println(c)
	c1 := Circle{0, 0, 5}
	fmt.Println("Second way to declare Struct", c1)
	c2 := new(Circle)
	*c2 = Circle{0, 0, 5}
	fmt.Println("Third way to declare Struct", c1)
	circleArea(c)
	// Method example
	fmt.Println("Method Example", c.area())
}
