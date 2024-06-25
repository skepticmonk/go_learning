package main

import "fmt"

func f() (int, int) {
	return 5, 6
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func function_learn() {
	// Function returning multiple Values
	x, y := f()
	fmt.Println("Function returning multiple Values: ", x, y)
	fmt.Println("Variadic function with multiple parameters: ", add(1, 2, 3))
	z := 0
	increment := func() int {
		z++
		return z
	}
	increment()
	fmt.Println("Closure: ", increment())
	fmt.Println("Recursion (Factorial of 10):", factorial(10))

}
