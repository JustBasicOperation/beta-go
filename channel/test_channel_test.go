package channel

import (
	"fmt"
	"sync"
	"testing"
)

func Test_channel01(t *testing.T) {
	ch := make(chan int, 1000)

	for i := 0; i < 1000; i++ {
		ch <- i
	}
	close(ch)

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for c := range ch {
				fmt.Printf("get int: %v\n", c)
			}
			wg.Done()
		}()
	}
	//time.Sleep(3*time.Second)
	wg.Wait()
}
