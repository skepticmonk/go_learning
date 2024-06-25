package channel

import (
	"fmt"
	"time"
)

func Select_example() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second): // Timeout
				fmt.Println("timeout")
				// default:
				// 	fmt.Println("nothing ready") // Default is implemented immediately,
				// if nothing is ready, so commenting out
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}
