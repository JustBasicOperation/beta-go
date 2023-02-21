package beta

// 闭包是由函数及其相关引用环境组合而成的实体，即闭包=函数+引用环境
// Go语言能通过escape analyze识别出变量的作用域，自动将变量在堆上分配。将闭包环境变量在堆上分配是Go实现闭包的基础。
// 返回闭包时并不是单纯返回一个函数，而是返回了一个结构体，记录下函数返回地址和引用的环境中的变量地址。
func closure(i int) func() int {
	return func() int {
		return i + 1
	}
}

func callClosure() {
	a1 := 1
	a2 := 2
	_ = closure(a1) // 闭包1
	_ = closure(a2) // 闭包2
}
