package p557

import "testing"

func TestReverseWords(t *testing.T) {
	input := "Let's take LeetCode contest"
	target := "s'teL ekat edoCteeL tsetnoc"
	if ReverseWords(input) != target {
		t.Error("reverse words in a string iii failed")
	}
}
