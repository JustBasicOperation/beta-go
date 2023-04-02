package beta

import (
	"fmt"
)

func testRedo() {
	fmt.Println("start redo")
	a := 0
redo:
	fmt.Println("begin redo")
	fmt.Println("middle redo")
	fmt.Println("end redo")
	a++
	if a == 1 {
		goto redo
	} else {
		fmt.Println("end")
		return
	}
}
