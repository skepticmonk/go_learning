package main

import (
	"fmt"
	"math/rand"
	"time"
)

func routine(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func go_routine() {
	fmt.Println("Go Routine")
	for i := 0; i < 10; i++ {
		go routine(i)
	}
	var input string
	_, err := fmt.Scan(&input)
	fmt.Println(err)
}
