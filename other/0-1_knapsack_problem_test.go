package other

import (
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func Test_maxWeightUsingDfs(t *testing.T) {
	type args struct {
		weight    []int
		maxWeight int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case-1",
			args: args{
				weight:    []int{2, 2, 4, 6, 3},
				maxWeight: 9,
			},
			want: 9,
		},
		{
			name: "case-2",
			args: args{
				weight:    []int{2, 2, 4, 6, 6},
				maxWeight: 9,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxWeightUsingDfs(tt.args.weight, tt.args.maxWeight); got != tt.want {
				t.Errorf("maxWeightUsingDfs() = %v, want %v", got, tt.want)
			}

			if got := maxWeightUsingMemoDfs(tt.args.weight, tt.args.maxWeight); got != tt.want {
				t.Errorf("maxWeightUsingMemoDfs() = %v, want %v", got, tt.want)
			}

			if got := maxWeightUsingDynamicPrograming(tt.args.weight, tt.args.maxWeight); got != tt.want {
				t.Errorf("maxWeightUsingDynamicPrograming() = %v, want %v", got, tt.want)
			}

			if got := maxWeightUsingDynamicProgramingWithLessMemory(tt.args.weight, tt.args.maxWeight); got != tt.want {
				t.Errorf("maxWeightUsingDynamicPrograming() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxValueUsingDynamicPrograming(t *testing.T) {
	type args struct {
		weight    []int
		value     []int
		maxWeight int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case-1",
			args: args{
				weight:    []int{2, 2, 4, 6, 3},
				value:     []int{3, 4, 8, 9, 6},
				maxWeight: 9,
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValueUsingDfs(tt.args.weight, tt.args.value, tt.args.maxWeight); got != tt.want {
				t.Errorf("maxValueUsingDfs() = %v, want %v", got, tt.want)
			}

			if got := maxValueUsingDynamicPrograming(tt.args.weight, tt.args.value, tt.args.maxWeight); got != tt.want {
				t.Errorf("maxValueUsingDynamicPrograming() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumLimitValueUsingDynamicPrograming(t *testing.T) {
	type args struct {
		value []int
		limit int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case-1",
			args: args{
				value: []int{10, 50, 30, 78, 93, 124, 25},
				limit: 200,
			},
			want: []int{93, 78, 30},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumLimitValueUsingDynamicPrograming(tt.args.value, tt.args.limit); !utils.ElementEqual(got, tt.want) {
				t.Errorf("minimumLimitValueUsingDynamicPrograming() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yanghuiTriangleUsingDynamicProgram(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want int
	}{
		{
			name: "case-1",
			args: [][]int{{5}, {7, 8}, {2, 3, 4}, {4, 9, 6, 1}, {2, 7, 9, 4, 5}},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := yanghuiTriangleUsingDynamicProgram(tt.args); got != tt.want {
				t.Errorf("yanghuiTriangleUsingDynamicProgram() = %v, want %v", got, tt.want)
			}
		})
	}
}
