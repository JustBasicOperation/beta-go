package beta

import (
	"fmt"
	"testing"
)

func Test_f1(t *testing.T) {
	i := f1()
	fmt.Printf("f1 return: %v\n", i)
}

func Test_f2(t *testing.T) {
	i := f2()
	fmt.Printf("f2 return: %v\n", i)
}

func Test_f3(t *testing.T) {
	i := f3()
	fmt.Printf("f3 return: %v\n", i)
}

func Test_f4(t *testing.T) {
	i := f4()
	fmt.Printf("f4 return: %v\n", i)
}

func Test_f8(t *testing.T) {
	f8()
}

func Test_f9(t *testing.T) {
	i := f9()
	fmt.Printf("f9 return: %v\n", i)
}
