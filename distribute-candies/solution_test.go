package p575

import "testing"

func TestDistributeCandies(t *testing.T) {
	candies := []int{1, 1, 2, 3}
	if 2 != DistributeCandies(candies) {
		t.Error("distribute candies failed")
	}
}
