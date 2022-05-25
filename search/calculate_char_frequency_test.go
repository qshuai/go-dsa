package search

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name                    string
	args                    string
	ExpectedCaseInsensitive []int
	ExpectedCaseSensitive   []int
}{
	{
		name:                    "empty string",
		args:                    "",
		ExpectedCaseInsensitive: nil,
		ExpectedCaseSensitive:   nil,
	},
	{
		name:                    "lowercase string",
		args:                    "helloworld",
		ExpectedCaseInsensitive: []int{0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 3, 0, 0, 2, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0},
		ExpectedCaseSensitive: []int{0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 3, 0, 0, 2, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		name:                    "uppercase string",
		args:                    "HELLOWORLD",
		ExpectedCaseInsensitive: []int{0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 3, 0, 0, 2, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0},
		ExpectedCaseSensitive: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 3, 0, 0, 2, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0},
	},
	{
		name:                    "mixed string",
		args:                    "hEllOworld",
		ExpectedCaseInsensitive: []int{0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 3, 0, 0, 2, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0},
		ExpectedCaseSensitive: []int{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 3, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0,
			0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
}

func Test_calculateCharFrequency(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcCharCntWithCaseInsensitive(tt.args); !reflect.DeepEqual(got, tt.ExpectedCaseInsensitive) {
				t.Errorf("CalcCharCntWithCaseInsensitive[%s] = %v, want %v", tt.name, got, tt.ExpectedCaseInsensitive)
			}

			if got := CalcCharCntWithCaseSensitive(tt.args); !reflect.DeepEqual(got, tt.ExpectedCaseSensitive) {
				t.Errorf("CalcCharCntWithCaseSensitive[%s] = %v, want %v", tt.name, got, tt.ExpectedCaseSensitive)
			}
		})
	}
}
