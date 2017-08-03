package p461

import "testing"

func TestHammingDistance(t *testing.T) {
	var x, y = 1, 4

	if 2 != HammingDistance(x, y) {
		t.Error("hamming distance failed")
	}
}
