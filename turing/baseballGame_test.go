package turring

import (
	"testing"
)

func TestCallPoints(t *testing.T) {
	tests := []struct {
		name string
		ops  []string
		want int
	}{
		{
			name: "Simple D Simple C Simple +",
			ops:  []string{"5", "2", "C", "D", "+"},
			want: 30,
		},
		{
			name: "Negative number duble +",
			ops:  []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			want: 27,
		},
		{
			name: "Single Score",
			ops:  []string{"1"},
			want: 1,
		},
		{
			name: "Single Sum",
			ops:  []string{"1", "1", "+"},
			want: 4,
		},
		{
			name: "Single Invalidate",
			ops:  []string{"1", "C"},
			want: 0,
		},
		{
			name: "Single Invalidate Two ops",
			ops:  []string{"3", "1", "C"},
			want: 3,
		},
		{
			name: "Single Double",
			ops:  []string{"1", "D"},
			want: 3,
		},
		{
			name: "No ops",
			ops:  []string{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := callPoints(tt.ops); got != tt.want {
				t.Errorf("callPoints(%s) = %d | want: %d", tt.ops, got, tt.want)
			}
		})
	}
}
