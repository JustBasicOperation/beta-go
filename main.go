package main

import (
	"fmt"
	"time"
)

//func main() {
//	s := []string{"aabc", "aac_d", "aaf"}
//	fmt.Printf("sorted: %v", s)
//	s1 := "f757b3d6bd1"
//	s2 := "f757b3d6bd"
//	fmt.Println(strings.Compare(s1, s2))
//}

type s1 struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type Value struct {
	ConfigValueType int32  `json:"config_value_type,omitempty"`
	ConfigValue     int32  `json:"config_value"`
	Value           string `json:"value,omitempty"`
	Other           int32  `json:"other"`
	Student         s1     `json:"student"`
}

func testTimeTransfer() {
	now := time.Now()
	utc := time.Now().UTC()
	local := time.Now().Local()
	zone, offset := time.Now().Zone()
	fmt.Println(now, utc, local, zone, offset)
}
