package basic_date_type

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_passSlice(t *testing.T) {
	s := []byte{'1'}
	fmt.Printf("address: %p, value: %v\n", &s, s)
	passSlice(s)
	fmt.Printf("address: %p, value: %v\n", &s, s)
}

func Test_passChannel(t *testing.T) {
	ch := make(chan byte, 2)
	fmt.Printf("address: %p, value: %v\n", &ch, ch)
	ch <- '1'
	passChannel(ch)
	fmt.Printf("address: %p, value: %v\n", &ch, ch)
}

func Test_constantEpx(t *testing.T) {
	// 整数(int)，rune，浮点数，复数，后面的可以兼容前面的类型
	// 隐式类型转换：15 + 25 + 5.2 -> int+int+float64 -> float64
	const constantEpx1 = 15 + 25 + 5.2
	// 隐式类型转换：15 + 25 + 5 -> int+int+int -> int
	const constantEpx2 = 15 + 25 + 5
	// 隐式类型转换：15 + 25 + int64(5) -> int+int+int64 -> int64
	const constantEpx3 = 15 + 25 + int64(5)
	// 隐式类型转换：15 + 25 + float64(5.6) -> int+int+float64 -> float64
	const constantEpx4 = 15 + 25 + float64(5.6)

	fmt.Println(constantEpx1, reflect.TypeOf(constantEpx1))
	fmt.Println(constantEpx2, reflect.TypeOf(constantEpx2))
	fmt.Println(constantEpx3, reflect.TypeOf(constantEpx3))
	fmt.Println(constantEpx4, reflect.TypeOf(constantEpx4))
}

func Test_callPassEnumType(t *testing.T) {
	callPassEnumType()
}

func Test_integerLiteral(t *testing.T) {
	integerLiteral()
}

func Test_rangeWithReplica(t *testing.T) {
	rangeWithReplica()
}
