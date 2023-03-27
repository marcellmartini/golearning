// You can edit this code!

// Click here and start typing.

package bubble

import (
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort()
		})
	}
}

func Test_bubleSort(t *testing.T) {
	type args struct {
		arr []int32
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Teste with 5 elements",
			args: args{
				[]int32{5, 1, 2, 4, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubleSort(tt.args.arr)
		})
	}
}
