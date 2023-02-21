package beta

import (
	"fmt"
	"testing"
	"time"
)

// 以下代码有什么问题，怎么解决？
func Test20230220_1(t *testing.T) {
	total, sum := 0, 0
	for i := 1; i <= 10; i++ {
		sum += i
		go func() {
			total += i
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Printf("total:%d sum %d\n", total, sum)
}
