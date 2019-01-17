package golang

import (
	"reflect"
	"testing"
)

func TestSplitStringWithRepeatedChar(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected []string
	}{
		{
			name:     "only one character",
			str:      "a",
			expected: []string{"a"},
		},
		{
			name:     "two different characters",
			str:      "ab",
			expected: []string{"ab"},
		},
		{
			name:     "two same characters",
			str:      "aa",
			expected: nil,
		},
		{
			name:     "none duplicated character",
			str:      "welcometomyhome",
			expected: []string{"welcometomyhome"},
		},
		{
			name:     "a duplicated character",
			str:      "helloworld",
			expected: []string{"he", "oworld"},
		},
		{
			name:     "start with duplicated character",
			str:      "aargwelldone",
			expected: []string{"rgwe", "done"},
		},
		{
			name:     "end with duplicated character",
			str:      "employeeexpress",
			expected: []string{"employ", "xpre"},
		},
	}

	for _, test := range tests {
		ret := SplitStringWithRepeatedChar(test.str)
		if !reflect.DeepEqual(ret, test.expected) {
			t.Errorf("%s, want: %v, but got: %v",
				test.name, test.expected, ret)

		}
	}
}
