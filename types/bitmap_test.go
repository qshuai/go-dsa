package types

import (
	"reflect"
	"strings"
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func TestBitMap_Add(t *testing.T) {
	tests := []struct {
		name  string
		b     *BitMap
		eles  []uint64
		want  []uint64
		panic bool
	}{
		{
			name:  "case-1",
			b:     NewBitMap(0, 65),
			eles:  []uint64{1, 64},
			want:  []uint64{1 << 62, 1 << 63},
			panic: false,
		},
		{
			name:  "case-2",
			b:     NewBitMap(0, 65),
			eles:  []uint64{65},
			want:  nil,
			panic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func() {
				for _, ele := range tt.eles {
					tt.b.Add(ele)
				}
			}

			utils.PanicHelper(t, fn, tt.panic)
			if !tt.panic && !reflect.DeepEqual(tt.b.data, tt.want) {
				t.Errorf("Add() = %v, want: %v", tt.b.data, tt.want)
			}
		})
	}
}

func TestBitMap_Contain(t *testing.T) {
	tests := []struct {
		name string
		b    *BitMap
		ele  uint64
		want bool
	}{
		{
			name: "caes-1",
			b: &BitMap{
				base: 0,
				bits: 65,
				data: []uint64{1 << 62, 0},
			},
			ele:  1,
			want: true,
		},
		{
			name: "caes-2",
			b: &BitMap{
				base: 0,
				bits: 65,
				data: []uint64{(1<<63 - 1) ^ (1 << 61), 1<<63 - 1},
			},
			ele:  2,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Contain(tt.ele); got != tt.want {
				t.Errorf("Contain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitMap_Remove(t *testing.T) {
	tests := []struct {
		name string
		b    *BitMap
		ele  uint64
		want []uint64
	}{
		{
			name: "case-1",
			b: &BitMap{
				base: 0,
				bits: 65,
				data: []uint64{1 << 63, 0},
			},
			ele:  0,
			want: []uint64{0, 0},
		},
		{
			name: "case-2",
			b: &BitMap{
				base: 0,
				bits: 65,
				data: []uint64{(1 << 62) | (1 << (63 - 34)), 0},
			},
			ele:  34,
			want: []uint64{1 << 62, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.b.Remove(tt.ele); !reflect.DeepEqual(tt.b.data, tt.want) {
				t.Errorf("Remove() = %v, want: %v", tt.b.data, tt.want)
			}
		})
	}
}

func TestBitMap_String(t *testing.T) {
	tests := []struct {
		name string
		b    *BitMap
		want string
	}{
		{
			name: "empty bits",
			b:    nil,
			want: "",
		},
		{
			name: "case-1",
			b:    NewBitMap(0, 64),
			want: strings.Repeat("0", 64),
		},
		{
			name: "case-2",
			b:    NewBitMap(0, 65),
			want: strings.Repeat("0", 65),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
