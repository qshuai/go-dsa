package leetcode

import (
	"reflect"
	"testing"
)

func Test_findErrorNums(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "case-1",
			nums: []int{1, 2, 2, 4},
			want: []int{2, 3},
		},
		{
			name: "case-2",
			nums: []int{1, 1},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findErrorNums(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findErrorNums() = %v, want %v", got, tt.want)
			}

			if got := findErrorNums2(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findErrorNums2() = %v, want %v", got, tt.want)
			}

			if got := findErrorNums3(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findErrorNums3() = %v, want %v", got, tt.want)
			}
		})
	}
}
