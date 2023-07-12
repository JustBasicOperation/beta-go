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

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func Test_app(t *testing.T) {
	a := app()
	b := app()
	a("go")
	fmt.Println(b("All"))
}
