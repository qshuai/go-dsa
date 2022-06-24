package other

import "testing"

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
