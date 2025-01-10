package utils

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestContain(t *testing.T) {
	type args[T comparable] struct {
		container []T
		target    T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}

	tests := []testCase[int]{
		{
			name: "contained",
			args: args[int]{
				container: []int{1, 2, 3},
				target:    1,
			},
			want: true,
		},
		{
			name: "empty container",
			args: args[int]{
				container: nil,
				target:    1,
			},
			want: false,
		},
		{
			name: "not contained",
			args: args[int]{
				container: []int{1, 2, 3},
				target:    0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contain(tt.args.container, tt.args.target); got != tt.want {
				t.Errorf("Contain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	type args[T any] struct {
		src []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}

	tests := []testCase[int]{
		{
			name: "empty slice",
			args: args[int]{
				src: nil,
			},
			want: nil,
		},
		{
			name: "copy",
			args: args[int]{
				src: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Copy(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElementEqual(t *testing.T) {
	type args[T comparable] struct {
		a []T
		b []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}

	tests := []testCase[int]{
		{
			name: "same sequence",
			args: args[int]{
				a: []int{2, 1, 3},
				b: []int{1, 2, 3},
			},
			want: true,
		},
		{
			name: "sequences with different length",
			args: args[int]{
				a: []int{1, 2, 3},
				b: []int{1, 2},
			},
			want: false,
		},
		{
			name: "same character set",
			args: args[int]{
				a: []int{1, 2, 2},
				b: []int{1, 1, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ElementEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ElementEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmbeddedSliceEqual(t *testing.T) {
	type args[T constraints.Ordered] struct {
		a [][]T
		b [][]T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want bool
	}

	tests := []testCase[int]{
		{
			name: "equaled",
			args: args[int]{
				a: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				b: [][]int{{4, 5, 6}, {1, 2, 3}, {7, 8, 9}},
			},
			want: true,
		},
		{
			name: "sequences with different length",
			args: args[int]{
				a: [][]int{{1, 2, 3}, {4, 5, 6}},
				b: [][]int{{4, 5, 6}},
			},
			want: false,
		},
		{
			name: "not equaled",
			args: args[int]{
				a: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				b: [][]int{{1, 2, 3}, {4, 5, 10}, {7, 8, 9}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EmbeddedSliceEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("EmbeddedSliceEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainEqual(t *testing.T) {
	type args[T constraints.Ordered] struct {
		arrs   [][]T
		target []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "case-1",
			args: args[int]{
				arrs:   [][]int{nil, {1, 2, 3}},
				target: nil,
			},
			want: true,
		},
		{
			name: "case-2",
			args: args[int]{
				arrs:   [][]int{{4, 5, 6}, {1, 2, 3}},
				target: nil,
			},
			want: false,
		},
		{
			name: "case-3",
			args: args[int]{
				arrs:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				target: []int{5, 4, 6},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainEqual(tt.args.arrs, tt.args.target); got != tt.want {
				t.Errorf("ContainEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
