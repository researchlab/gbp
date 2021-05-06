package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	r := Fib(5)
	if r != 5 {
		t.Error("Fib(5) is not equal to 5")
	}
}
