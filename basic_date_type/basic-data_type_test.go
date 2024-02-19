package basic_date_type

import (
	"fmt"
	"reflect"
	"testing"
	"unicode/utf8"
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

func Test_constantExp(t *testing.T) {
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

func Test_varExp(t *testing.T) {
	var a = 0
	fmt.Println(reflect.TypeOf(a))
	b := returnFloat64()
	fmt.Println(reflect.TypeOf(b))
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

func Test_testSlice01(t *testing.T) {
	testSlice01()
}

func TestRune01(t *testing.T) {
	s1 := "Go语言程序1234"
	bytes := []byte(s1)
	runeBytes := []rune(s1)
	fmt.Printf("原始字符串：%v\n", s1)
	fmt.Printf("对应的二进制字节序列：%v\n", bytes)
	fmt.Printf("对应的rune序列：%v\nUnicode表示: %U\n", runeBytes, runeBytes)
}

func TestRune02(t *testing.T) {
	s1 := "语"
	bytes := []byte(s1)
	r, size := utf8.DecodeRune(bytes)
	fmt.Println(r, size)
}

// 可以通过s[0]的方式访问s中的字节，但是不能改变s[0]的值
func TestStringType(t *testing.T) {
	s := "123"
	//s[0] = '2'
	fmt.Println(s[0])

	// 字节数组的元素取地址是合法的，但是对字符串字节数组的元素取地址是非法的
	b := []byte{1, 2, 3}
	p := &b[0] // 合法
	fmt.Println(p)

	//p1 := &s[0] 非法
}

func TestArrayAndSlice(t *testing.T) {
	var a [10]int // 数组在声明时会分配内存并初始化零值
	var s []int   // 切片声明后没有赋值时 == nil
	fmt.Println(a, s == nil)
	fmt.Printf("a: %p, s:%p\n", &a, &s)

	var a1 = new([10]int)
	var s1 = new([]int)
	fmt.Println(a1 == nil, s1 == nil)
	fmt.Println(a1, s1)
}

func TestSlice20230726(t *testing.T) {
	//s := []string{"1", "2"}
	s1 := make([]string, 100, 100)
	//fmt.Printf("s99: %v\n", s[99]) // panic: runtime error: index out of range [99] with length 2
	fmt.Printf("s199: %v\n", s1[0]) // 空字符串
}

func Test_testPassIntSlice(t *testing.T) {
	testPassIntSlice()
}
