package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	borad := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	// Given word = "ABCCED", return true.
	// Given word = "SEE", return true.
	// Given word = "ABCB", return false.
	//t.Log("ABCCED=", exist(borad, "ABCCED"))
	t.Log("SEE=", exist2(borad, "SEE"))
	//t.Log("ABCB=", exist(borad, "ABCB"))
}

func Test_exist2(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				board: [][]byte{
					{'C', 'A', 'A'},
					{'A', 'A', 'A'},
					{'B', 'C', 'D'},
				},
				word: "AAB",
			},
			want: true,
		},
		{
			name: "t2",
			args: args{
				[][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'E', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				"ABCEFSADEESE",
			},
			want: true,
		},
		{
			name: "t3",
			args: args{
				board: [][]byte{
					{'a', 'a', 'a', 'a'},
					{'a', 'a', 'a', 'a'},
					{'a', 'a', 'a', 'a'},
				},
				word: "aaaaaaaaaaaaa",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist2(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist2() = %v, want %v", got, tt.want)
			}
		})
	}
}
