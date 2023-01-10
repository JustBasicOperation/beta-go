package beta

import (
	"fmt"
)

func testDefer() {
	defer func() {}()
}

/*
前置知识:
pre1: return xxx语句不是原子的，return xxx语句分为两步，
第一步是给返回值赋值(如果是返回值是匿名的，则先声明返回值变量，再给返回值变量赋值)，如果xxx是函数，则先执行函数
第二步是执行ret指令
pre2: panic和recover是内置的函数，本质上也是函数

defer知识点：
1.多个defer的执行顺序
  先进后出
2.defer的触发时机，遇见return，或者遇见panic
3.defer的执行时机
  遇见return：在return语句的中间执行；遇见panic：在panic函数之后执行，如果panic之后同时有recover和defer，则优先执行recover
4.defer panic 和 recover的执行情况，case见f5和f6
  4.1 panic语句之后的defer语句永远不会被执行
  4.2 panic之后，遍历本协程的defer栈并执行defer，如果遇到recover则返回recover处继续执行；如果没有遇到recover则执行完defer栈后抛出报错信息
5.defer中有panic的情况，case见f7
  这种情况相当于有多个panic，则只有最后一个panic会被recover捕获，前面执行的panic会覆盖后面执行的panic
*/

// 练习 case
// f1
// result: 1，
// 解析: 具名变量i会在函数起始处被初始化为零值
func f1() (i int) {
	defer func() {
		i++
	}()
	return 0
}

// f2
// result: 5
func f2() (r int) {
	i := 5
	defer func() {
		i += 5
	}()
	return i
}

// f3
// result: 1，
// 解析: defer函数里面的r变量和外面的r变量作用域不同，go中都是值传递，所以返回值为1
func f3() (r int) {
	r = 1
	defer func(r int) {
		r += 5
	}(r)
	return r
}

// f4
// result: 1
func f4() int {
	i := 1
	defer func() {
		i += 10
	}()
	return i
}

func f5() {
	defer func() { fmt.Println("defer: panic 之前1") }()
	defer func() { fmt.Println("defer: panic 之前2") }()

	panic("异常内容") //触发defer出栈

	defer func() { fmt.Println("defer: panic 之后，永远执行不到") }()
}

func f6() {
	defer func() {
		fmt.Println("defer: panic 之前1, 捕获异常")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() { fmt.Println("defer: panic 之前2, 不捕获") }()

	panic("异常内容") //触发defer出栈

	defer func() { fmt.Println("defer: panic 之后, 永远执行不到") }()
}

func f7() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("panic")
}

func function(index int, value int) int {
	fmt.Println(index)
	return index
}

// result: 3 4 2 1
// 解析: 这里面有两个defer， 所以defer一共会压栈两次，先进栈1，后进栈2。
// 在压栈function1的时候，需要连同函数地址、函数形参一同进栈，为了得到function1的第二个参数的结果，
// 所以就需要先执行function3将第二个参数算出，所以function3就被第一个执行。
// 同理压栈function2，就需要执行function4算出function2第二个参数的值。
// 然后函数结束，先出栈function2、再出栈function1.
// 所以执行顺序如下：
// defer压栈function1，压栈函数地址、形参1、形参2(调用function3) --> 打印3
// defer压栈function2，压栈函数地址、形参1、形参2(调用function4) --> 打印4
// defer出栈function2, 调用function2 --> 打印2
// defer出栈function1, 调用function1--> 打印1
func f8() {
	defer function(1, function(3, 0))
	defer function(2, function(4, 0))
}

// f9: 面试易错题
// result: 2
// 执行顺序：
// 1.初始化返回值t为零值 0
// 2.首先执行defer的第一步，赋值defer中的func入参t为0
// 3.执行defer的第二步，将defer压栈
// 4.将t赋值为1
// 5.执行return语句，将返回值t赋值为2
// 6.执行defer的第三步，出栈并执行
// 7.因为在入栈时defer执行的func的入参已经赋值了，此时它作为的是一个形式参数，所以打印为0；
// 相对应的因为最后已经将t的值修改为2，所以再打印一个2
func f9() (t int) {
	defer func(i int) {
		fmt.Println(i)
		fmt.Println(t)
	}(t)
	t = 1
	return 2
}
