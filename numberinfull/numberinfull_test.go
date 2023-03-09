package numberinfull

import (
	"testing"
)

func TestNumberINFull(t *testing.T) {
	tests := []struct {
		name   string
		number int32
		want   string
	}{
		{
			"Number in full 1111",
			1111,
			"mil cento e onze",
		},
		{
			"Number in full 2400",
			2400,
			"dois mil e quatrocentos",
		},
		{
			"Number in full 18871",
			18871,
			"dezoito mil oitocentos e setenta e um",
		},
		{
			"Number in full 50100",
			50100,
			"cinquenta mil e cem",
		},
		{
			"Number in full 78383",
			78383,
			"setenta e oito mil trezentos e oitenta e tres",
		},
		{
			"Number in full 1234323456",
			1234323456,
			"numero fora do limite",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberINFull(tt.number); got != tt.want {
				t.Errorf("numberINFull(%v) = %v, want: %v",
					tt.number, got, tt.want)
			}
		})
	}
}
