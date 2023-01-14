package main

import (
	"nu/utils"
	"testing"
)

type args struct {
	pontuacao     []int32
	tamanhoDoTime int32
	k             int32
}

type test struct {
	name string
	args args
	want int64
}

func BenchmarkFormacaoDeTime(b *testing.B) {
	test := test{
		args: args{
			utils.GetPontuacao3(),
			utils.GetTamanhoTime3(),
			utils.GetK3(),
		},
		want: utils.GetResult3(),
	}

	for i := 0; i < b.N; i++ {
		_ = formacaoDeTime(test.args.pontuacao, test.args.tamanhoDoTime, test.args.k)
	}
}

func TestFormacaoDeTime(t *testing.T) {
	tests := []test{
		{
			name: "Time1",
			args: args{
				utils.GetPontuacao1(),
				utils.GetTamanhoTime1(),
				utils.GetK1(),
			},
			want: utils.GetResult1(),
		}, {
			name: "Time2",
			args: args{
				utils.GetPontuacao2(),
				utils.GetTamanhoTime2(),
				utils.GetK2(),
			},
			want: utils.GetResult2(),
		}, {
			name: "Time3",
			args: args{
				utils.GetPontuacao3(),
				utils.GetTamanhoTime3(),
				utils.GetK3(),
			},
			want: utils.GetResult3(),
		}, {
			name: "Time4",
			args: args{
				[]int32{1, 2, 3, 4},
				3,
				1,
			},
			want: 9,
		}, {
			name: "Time5",
			args: args{
				[]int32{1},
				1,
				1,
			},
			want: 1,
		}, {
			name: "Time6",
			args: args{
				[]int32{1, 2},
				1,
				1,
			},
			want: 2,
		}, {
			name: "Time7",
			args: args{
				[]int32{1, 2, 3},
				1,
				2,
			},
			want: 3,
		}, {
			name: "Time8",
			args: args{
				[]int32{15, 15, 15, 12},
				2,
				1,
			},
			want: 30,
		}, {
			name: "Time9",
			args: args{
				[]int32{15, 15, 15, 12},
				3,
				2,
			},
			want: 45,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formacaoDeTime(tt.args.pontuacao, tt.args.tamanhoDoTime, tt.args.k); got != tt.want {
				t.Errorf("formacaoDeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMain(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "main"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
