package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func pointers() {
	x := 5
	zero(&x)
	fmt.Println(x)
	xPtr := new(int)
	one(xPtr)
	fmt.Println("Pointer with new operator", *xPtr)
}
