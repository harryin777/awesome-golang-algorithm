package Solution

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []string
		expect string
	}{
		{"TestCase", []string{"ADOBECODEBANC", "ABC"}, "BANC"},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			//got := minWindow(c.inputs[0], c.inputs[1])
			//if !reflect.DeepEqual(got, c.expect) {
			//	t.Fatalf("expected: %v, but got: %v, with inputs: %v",
			//		c.expect, got, c.inputs)
			//}
		})
	}
}

func Test_minWindow4(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				"ADOBECODEBANC",
				"ABC",
			},
			want: "BANC",
		},
		{
			name: "t2",
			args: args{
				"a",
				"aa",
			},
			want: "",
		},
		{
			name: "t2",
			args: args{
				"a",
				"a",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow4(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_4(t *testing.T) {
	fmt.Println(minWindow4("ADOBECODEBANC", "ABC"))
}
