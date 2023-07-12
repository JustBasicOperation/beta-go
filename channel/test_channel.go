package channel

import (
	"fmt"
	"strconv"
	"time"
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

func talk(msg string, sleep int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(sleep) * time.Millisecond)
		}
	}()
	return ch
}

func fanIn(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case a := <-input1:
				ch <- a
			case b := <-input2:
				ch <- b
			}
		}
	}()
	return ch
}

func test01() {
	ch := fanIn(talk("A", 10), talk("B", 1000))
	for range ch {
		fmt.Printf("%q\n", <-ch)
	}
}

var ch = make(chan string, 10)

func produce() {
	defer func() {
		close(ch)
	}()
	time.Sleep(3 * time.Second)
	for i := 0; i < 100; i++ {
		ch <- strconv.FormatInt(int64(i), 10)
	}
}

func consumer() {
	for s := range ch {
		fmt.Printf("consumer ch: %v\n", s)
	}
}
