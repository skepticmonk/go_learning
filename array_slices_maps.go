package main

import "fmt"

func array_slices_maps() {
	var y [5]int
	y[4] = 100
	fmt.Println(y)
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
	// Create slice using make functions
	z := make([]float64, 5)
	fmt.Println(z)
	arr := [5]float64{1, 2, 3, 4, 5}
	a := arr[0:5]
	fmt.Println(a)

	// Slice functions
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	// Maps
	var m map[string]int
	m = make(map[string]int) // Another way to create map
	m["key"] = 10
	m["key2"] = 20
	for k, v := range m {
		fmt.Printf("key[%s] value[%d]\n", k, v)
	}
	delete(m, "key")
	fmt.Println(m)
}
