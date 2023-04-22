package goroutine

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan int, 10)

func testForSelect() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			readFromChan()
		}()
	}
	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func readFromChan() {
	for {
		select {
		case a := <-ch:
			fmt.Println(a)
		default:
			return
		}
	}
}
