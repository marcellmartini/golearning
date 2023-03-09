package main

import (
	"bytes"
	"testing"
)

func TestM5M3(t *testing.T) {
	tests := []struct {
		name   string
		number int
		want   string
	}{
		{
			name:   "Print number",
			number: 1,
			want:   "1\n",
		},
		{
			name:   "Print M3",
			number: 3,
			want:   "M3\n",
		},
		{
			name:   "Print M5",
			number: 5,
			want:   "M5\n",
		},
		{
			name:   "Print M3M5",
			number: 15,
			want:   "M3M5\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer := bytes.Buffer{}
			m3m5(&buffer, tt.number)

			if got := buffer.String(); got != tt.want {
				t.Errorf("m3m5(%v) = %v, want: %v", tt.number, got, tt.want)
			}
		})

	}
}
