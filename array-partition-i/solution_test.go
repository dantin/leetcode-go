package p561

import "testing"

func TestArrayPairSum(t *testing.T) {
	nums := []int{1, 4, 3, 2}

	if 4 != ArrayPairSum(nums) {
		t.Error("array partition i failed")
	}
}
