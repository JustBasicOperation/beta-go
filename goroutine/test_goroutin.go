package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func testForSelect() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			readFromChan(ch) // 启动十个协程轮询channel
		}()
	}
	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	wg.Wait()
}

func readFromChan(ch chan int) {
	for {
		select {
		case a, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(a)
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func printNum(n int) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		for num := range ch1 {
			if num > n {
				close(ch2)
				close(ch3)
				return
			}
			fmt.Printf("Goroutine 1 num: %v\n", num)
			ch2 <- num + 1
		}
	}()
	go func() {
		defer wg.Done()
		for num := range ch2 {
			if num > n {
				close(ch1)
				close(ch3)
				return
			}
			fmt.Printf("Goroutine 1 num: %v\n", num)
			ch3 <- num + 1
		}
	}()
	go func() {
		defer wg.Done()
		for num := range ch3 {
			if num > n {
				close(ch2)
				close(ch1)
				return
			}
			fmt.Printf("Goroutine 1 num: %v\n", num)
			ch1 <- num + 1
		}
	}()

	go func() {
		ch1 <- 1
	}()
	wg.Wait()
}
