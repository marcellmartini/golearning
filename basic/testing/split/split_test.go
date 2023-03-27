package split

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		name string
		s    string
		sep  string
		want []string
	}{
		{
			name: "simple",
			s:    "a/b/c",
			sep:  "/",
			want: []string{"a", "b", "c"},
		},
		{
			name: "wrong",
			s:    "a/b/c",
			sep:  ",",
			want: []string{"a/b/c"},
		},
		{
			name: "no sep",
			s:    "abc",
			sep:  "/",
			want: []string{"abc"},
		},
		{
			name: "trailing sep",
			s:    "a/b/c/",
			sep:  "/",
			want: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Split(tt.s, tt.sep)
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
