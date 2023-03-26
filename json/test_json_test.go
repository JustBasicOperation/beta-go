package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type MyPerson struct {
	name  string // 不可导出字段，json序列化时不会生效
	hobby string // 不可导出字段，json序列化时不会生效
}

func Test01(t *testing.T) {
	person := MyPerson{name: "polarisxu", hobby: "Golang"}
	m, _ := json.Marshal(person)
	fmt.Printf("%s\n", m)
}
