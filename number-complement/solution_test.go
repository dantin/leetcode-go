package p476

import "testing"

func TestFindComplement(t *testing.T) {
	num := 5
	if 2 != FindComplement(num) {
		t.Error("number complement error")
	}
}
