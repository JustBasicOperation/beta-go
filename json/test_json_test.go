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

type OriPerson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ModifyPerson struct {
	NewName string `json:"new_name"` // json反序列化的时候只会解析名称匹配的字段，忽略名称不匹配的字段
	Age     int    `json:"age"`
}

func Test02(t *testing.T) {
	p1 := &OriPerson{
		Name: "123",
		Age:  18,
	}
	p2 := &ModifyPerson{}
	marshal, _ := json.Marshal(p1)
	fmt.Printf("p1: \n%v\n", string(marshal))

	if err := json.Unmarshal(marshal, p2); err != nil {
		fmt.Printf("Unmarshal err: %v", err)
		return
	}
	fmt.Printf("p2: \n%v\n", p2)
}

func Test03(t *testing.T) {
	//p := &Person{
	//	FirstName:  "12",
	//	SecondName: "34",
	//}
	s := "{\"Name\":\"a\", \"Person\":{\"firstName\":\"1\",\"secondName\":\"2\"}}"
	h := &Human{}
	if err := json.Unmarshal([]byte(s), h); err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Printf("h: %v\n", h)
	fmt.Printf("h: %v\n", h.Person)
}
