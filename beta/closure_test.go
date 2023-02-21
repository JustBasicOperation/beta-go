package beta

import (
	"fmt"
	"testing"
)

func Test_closure(t *testing.T) {
	for i := 0; i < 10; i++ {
		f := closure(i)
		fmt.Println(f())
	}
}
