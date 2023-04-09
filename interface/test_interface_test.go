package _interface

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_testInterface(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testInterface()
		})
	}
}

func Test_testEmbeddedInterface(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testEmbeddedInterface()
		})
	}
}

func Test_01(t *testing.T) {
	var s *Student
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	var p People = s
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Printf("%v\n", reflect.TypeOf(p))
		fmt.Println(p)
		fmt.Println("p is not nil")
	}
}
