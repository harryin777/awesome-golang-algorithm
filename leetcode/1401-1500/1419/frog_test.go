package _419

import "testing"

func Test_minNumberOfFrogs(t *testing.T) {
	type args struct {
		croakOfFrogs string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				"crcoakroak",
			},
			want: 2,
		},
		{
			name: "t2",
			args: args{
				"croakcrook",
			},
			want: -1,
		},
		{
			name: "t3",
			args: args{
				"croakcroak",
			},
			want: 1,
		},
		{
			name: "t4",
			args: args{
				"croakroak",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumberOfFrogs(tt.args.croakOfFrogs); got != tt.want {
				t.Errorf("minNumberOfFrogs() = %v, want %v", got, tt.want)
			}
		})
	}
}
