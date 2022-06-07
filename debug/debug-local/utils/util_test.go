package utils

import "testing"

func TestAdd(t *testing.T) {
	if Add(100, 200) != 300 {
		t.Error("100 + 200 != 300")
	}
}
