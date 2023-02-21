package _interface

import (
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
