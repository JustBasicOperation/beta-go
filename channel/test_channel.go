package channel

import (
	"fmt"
)

func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
		fmt.Println("send date: ", i)
	}
}

func recv(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}
