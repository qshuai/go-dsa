package other

import (
	"testing"
)

func Test_backToOrigin(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "case-1",
			args: 2,
			want: 2, // 0->1->0, 0->9->0
		},
		{
			name: "case-2",
			args: 4,
			want: 6, // 0->1->2->1->0, 0->1->0->9->0, 0->1->0->1->0 相反方向也有3种情况
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backToOrigin(tt.args); got != tt.want {
				t.Errorf("backToOrigin() = %v, want %v", got, tt.want)
			}
		})
	}
}
