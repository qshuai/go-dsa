package count

import (
	"reflect"
	"testing"
)

func TestCalculateCharFrequency(t *testing.T) {
	str := "helloworld"

	expected := []int{0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 3,
		0, 0, 2, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0}

	if !reflect.DeepEqual(calculateCharFrequency(str), expected) {
		t.Errorf("error: want: %v, but got: %v", expected,
			calculateCharFrequency(str))
	}
}
