package channel

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func Channel_example() {
	var c chan string = make(chan string)
	//c := make(chan int, 1) ==> This would create a buffered channel,
	// which is asynchronous with capacity of 1.
	go pinger(c)
	go printer(c)
	var input string
	fmt.Scan(&input)

	// we can also specify the directions for these channels
	// func pinger(c chan<- string) ==> This would restrict pinger to sending
	// func printer(c <-chan string) ==> This would restrict printer to receiving
}
