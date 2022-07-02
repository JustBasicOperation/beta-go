package goroutine

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1e9)
}
