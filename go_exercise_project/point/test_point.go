package main

import (
	"fmt"
	"reflect"
	"time"
)

// 结论：
// 1.切片：使用make和var两种方式来测试，证明go中的切片不是引用类型，但是可以通过创建指针类型的切片来传递切片的引用
// 2.map：是引用类型，可以通过传递引用来操作map
// 3.chan：是引用类型，但是不能在一个线程中读写同一个channel，会产生死锁报错，必须通过协程发送数据或者读取数据

func main() {
	// 测试切片是否为引用类型
	// 使用make创建的切片
	s := make([]string, 0, 10)
	s1 := "1"
	s = append(s, s1)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(&s))
	Func1(&s)
	fmt.Println(s[0], s[1])

	// 使用var创建的切片
	var s2 []string
	s2 = append(s2, "1")
	Func3(s)
	fmt.Println(s2)

	// 测试map是否为引用类型
	m := make(map[string]string)
	fmt.Println(reflect.TypeOf(m))
	m["1"] = "2"
	Func2(m)
	fmt.Println(m)

	// 测试通道是否为引用类型
	c := make(chan int, 10)
	c <- 1
	go send(c)
	go receive(c)
	time.Sleep(3 * time.Second)

}

func Func1(s *[]string) {
	s2 := "123"
	*s = append(*s, s2)
}

func Func2(m map[string]string) {
	m["2"] = "3"
}

func Func3(s []string) {
	s = append(s, "321")
}

func receive(c chan int) {
	for i := range c {
		fmt.Println(i)
	}
}

func send(c chan int) {
	c <- 2
}
