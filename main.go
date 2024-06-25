package main

import (
	"fmt"

	"github.com/skepticmonk/go_learning/channel"
)

// Go does not have classes
// this is a comment
func main() {
	numberOfBits, errorMsg := fmt.Println("Hello World")
	fmt.Println(numberOfBits, errorMsg)
	n := 13
	// decalaring variable without specifying type,
	//don't use var when declaring variable without type
	//type is inferred and if we uncomment the below line, it will throw a compilation error
	// n = "Hello"
	fmt.Println(n)
	var k string
	k = "Let it be"
	fmt.Println(k)
	const o string = "Oh god"
	fmt.Println("What a life,", o)
	var r string
	fmt.Println("Enter your name:")
	fmt.Scanf("%s", &r)
	fmt.Println("Hello", r)
	control()
	array_slices_maps()
	function_learn()
	defer_panic_recover()
	pointers()
	struct_example()
	interface_example()
	go_routine()
	channel.Select_example()
	channel.Channel_example()
}
