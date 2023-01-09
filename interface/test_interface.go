package _interface

import (
	"fmt"
)

type A struct {
	a int
}

func (a A) f() {
	fmt.Println(a.a)
}

func (a A) f1() {
	fmt.Println(a.a)
}

type B struct {
	b int
}

func (b B) f() {
	fmt.Println(b.b)
}

type C int

func (c C) f() {
	fmt.Println(c)
}

type I interface {
	f()
}

type I1 interface {
	f1()
}

func testInterface() {
	// 1.接口是一种数据类型，可以接受任何类型的赋值
	//var i interface{} = 9
	//i = 1
	//i = "adc"
	//of := reflect.TypeOf(i)
	//fmt.Println(i)
	//fmt.Println(of)

	// 2.只要类型实现了接口的所有方法就能用接口的变量来接收该类型的变量
	// 如果类型实现了两个接口的所有的方法，那么也可以用这两个接口的变量来接收该类型的变量
	var a = A{a: 1}
	a.f()
	var b I = a
	b.f()

	var a1 = A{a: 2}
	a1.f()
	a1.f1()
	var b1 I1 = a1
	b1.f1()
}
