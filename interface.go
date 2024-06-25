package main

import "fmt"

type Shape interface {
	area() int64
}

type Square struct {
	side int64
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (s Square) area() int64 {
	return (s.side * s.side)
}

func interface_example() {
	var i Shape = Square{2}
	fmt.Println("Interface Example: ", i.area())
}
