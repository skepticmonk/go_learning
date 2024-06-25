package main

import "fmt"

// Go has a special statement called defer which sched
// ules a function call to be run after the function com
// pletes.
// In the below example as we can see, if we use defer,
// open and close functions are nearer to one another
// f, _ := os.Open(filename)
//  defer f.Close()

func wrong_way_to_implement_recover() {
	panic("PANIC")
	str := recover()
	fmt.Println(str)
}

func correct_way_to_implement_recover() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func defer_panic_recover() {
	defer second()
	first()
	correct_way_to_implement_recover()
	// wrong_way_to_implement_recover()
}
