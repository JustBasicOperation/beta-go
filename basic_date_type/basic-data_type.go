package basic_date_type

import (
	"fmt"
	"reflect"
)

// go语言基本数据类型
//func main() {
// 1.布尔类型
//var b bool = true
//fmt.Println(b)

// 2.数字类型
//var a1 uint8 = 1
//var a2 uint16 = 2
//var a3 uint32 = 3
//var a4 uint64 = 4
//var a5 int8 = 1
//var a6 int16 = 2
//var a7 int32 = 3
//var a8 int64 = 4
//fmt.Println(a1,a2,a3,a4,a5,a6,a7,a8)

// 3.浮点型
//var f1 float32 = 1.0 //精确到小数点后6位
//var f2 float64 = 2.0 //精确到小数点后15位
//fmt.Println(f1, f2)

// 4.其他数字类型
//var b1 byte = 'a'
//var b2 rune = 3
//var b3 int = 4
//var b4 uintptr = 5
//var b5 uint = 64 //uint和int的长度取决于操作系统的字长，32位系统就是32位，64位系统就是64位
//var b6 int = 64
//fmt.Println(b1,b2,b3,b4)

// 5.字符串
//var str string = "123"
//fmt.Println(strconv.Atoi(str)) //字符串转int
//fmt.Println(strconv.Itoa(456)) //int转字符串
//fmt.Println(strconv.ParseInt(str,10,64)) //
//fmt.Println(strconv.FormatInt(234,10))

// 6.unicode
//var ch int = '\u0041'
//var ch2 int = '\u03B2'
//var ch3 int = '\U00101234'
//fmt.Println(ch,ch2,ch3)

// 7.变量
//var 变量1 string = "变量1"
//fmt.Println(变量1)

// 8.变量零值
//int 为 0
//float 为 0.0
//bool 为 false
//string 为空字符串""
//指针为 nil
//nil 标志符用于表示interface、函数、maps、slices、channels、error、指针等的“零值”。
//var m map[string]int
//m["one"] = 1 //error  在一个 nil 的slice中添加元素是没问题的，但对一个map做同样的事将会生成一个运行时的panic

// 9.常量 定义：const identifier [type] = value
// Go语言预定义了这些常量： true、 false和iota,itoa在每一个const关键字出现时被重置为0，然后在下一个const出现之前，其所代表的数字会自动增1
//const (
//	a = iota
//	b = iota
//	c = iota
//)
//fmt.Println(a,b,c)

//const(
//	d = iota
//	e
//	f
//)
//fmt.Println(d,e,f) //后面的e，f没有定义，会自动继承上一个常量的值

// 10.作用域
// 有花括号"{ }"一般都存在作用域的划分；
// := 简式声明会屏蔽所有上层代码块中的变量（常量），建议使用规则来规范，如对常量使用全部大写，而变量尽量小写；
// 在if等语句中存在隐式代码块，需要注意；
//闭包函数可以理解为一个代码块，并且他可使用包含它的函数内的变量；
//if a := 1; false {
//} else if b := 2; false {
//} else if c := 3; false {
//} else {
//	println(a, b, c)
//}
//}

// 11.结构体
type student struct {
	name     string
	age      int32
	relation map[string]string
}

func testStruct() {
	s1 := &student{name: "name", relation: map[string]string{}}
	s2 := &student{name: "name", relation: nil}
	equal := reflect.DeepEqual(s1, s2)
	fmt.Println(equal)
}

// 12. go语言传值: 值传递
// 为什么说go语言的参数传递都是值传递？
// 1.对于切片，channel而言，传参时会复制一个切片或者channel传递，
//
//	也就是说传参前的和传参后是两个不同的切片或者channel，但是切片或channel指向的底层数组是一样的
//
// 2.对于指针类型，传参时会复制指针的值进行传递
func passSlice(s1 []byte) {
	fmt.Printf("paasSlice address: %p, value: %v\n", &s1, s1)
	s1[0] = '2'
	fmt.Printf("paasSlice address: %p, value: %v\n", &s1, s1)
}

func passChannel(ch chan byte) {
	fmt.Printf("paasSlice address: %p, value: %v\n", &ch, ch)
	ch <- '2'
	fmt.Printf("paasSlice address: %p, value: %v\n", &ch, ch)
}

// 3.补充：range是副本参与循环
func rangeWithReplica() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

// 13. go枚举类型传参
type intEnum int32
type intEnum2 intEnum

type strEnum string

const a intEnum = 1
const b intEnum = 2

func passEnumType(enum intEnum) {
	fmt.Println(enum)
}

func passEnumType2(enum intEnum2) {
	fmt.Println(enum)
}

func passEnumType3(enum strEnum) {
	fmt.Println(enum)
}

func callPassEnumType() {
	passEnumType(1)
	passEnumType2(2)
	passEnumType3("123")
}

/*
Integer literals

valid case:
42
4_2
0600
0_600
0o600
0O600       // second character is capital letter 'O'
0xBadFace
0xBad_Face
0x_67_7a_2f_cc_40_c6
170141183460469231731687303715884105727
170_141183_460469_231731_687303_715884_105727

invalid case:
_42         // an identifier, not an integer literal
42_         // invalid: _ must separate successive digits
4__2        // invalid: only one _ at a time
0_xBadFace  // invalid: _ must separate successive digits
*/

func integerLiteral() {
	fmt.Println(42)
	fmt.Println(4_2)
	fmt.Println(0600)
	fmt.Println(0_600)
	fmt.Println(0o600)
	fmt.Println(0o600)
	fmt.Println(0xBadFace)
	fmt.Println(0xBad_Face)
	fmt.Println(0xBadFace)
	fmt.Println(0x_67_7a_2f_cc_40_c6)
	var a float64 = 170141183460469231731687303715884105727
	fmt.Printf("%v\n", a)
	var b float64 = 170_141183_460469_231731_687303_715884_105727
	fmt.Printf("%v\n", b)
}

/*
测试slice语法糖
...slice传过去的是同一个切片，并不会生成一个新的切片
*/
func testSlice01() {
	s := make([]int, 5, 5)
	s[0] = 1
	s[1] = 2

	fmt.Printf("len: %d  cap: %d pointer: %p\n", len(s), cap(s), s)
	change(s...) // 这种形式传参，不会创建新的切片
	fmt.Println(s)
	fmt.Printf("len: %d  cap: %d pointer: %p\n", len(s), cap(s), s)
	fmt.Println("-------------------------")
	fmt.Printf("len: %d  cap: %d pointer: %p\n", len(s), cap(s), s)
	change(s[0:2]...) // 传入的还是原来的切片地址，不会创建新的切片
	fmt.Println(s)
	fmt.Printf("len: %d  cap: %d pointer: %p\n", len(s), cap(s), s)
}

func change(s ...int) {
	fmt.Printf("len: %d  cap: %d pointer: %p\n", len(s), cap(s), s)
	s = append(s, 3)
	fmt.Printf("len: %d  cap: %d pointer: %p\n", len(s), cap(s), s)
}
