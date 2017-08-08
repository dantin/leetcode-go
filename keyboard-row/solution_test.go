package p500

import (
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	input := []string{"Hello", "Alaska", "Dad", "Peace"}
	target := []string{"Alaska", "Dad"}
	if !reflect.DeepEqual(target, FindWords(input)) {
		t.Error("keyboard row failed")
	}
}
