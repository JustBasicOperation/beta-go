package basic_date_type

import (
	"fmt"
	"testing"
)

func Test_paasSlice(t *testing.T) {
	s := []byte{'1'}
	fmt.Printf("address: %p, value: %v\n", &s, s)
	paasSlice(s)
	fmt.Printf("address: %p, value: %v\n", &s, s)
}

func Test_paasChannel(t *testing.T) {
	ch := make(chan byte, 2)
	fmt.Printf("address: %p, value: %v\n", &ch, ch)
	ch <- '1'
	paasChannel(ch)
	fmt.Printf("address: %p, value: %v\n", &ch, ch)
}
