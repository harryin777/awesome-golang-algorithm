package Solution

import (
	"fmt"
	"testing"
)

func Test_wordBreakBFS(t *testing.T) {
	//fmt.Println(wordBreakBFS("aaaaaaaaaab", []string{"a", "aa", "aaa"}))
	//fmt.Println(wordBreak3("leetcode", []string{"leet", "code"}))
	//fmt.Println(wordBreak4("aaaaaaa", []string{"aaaa", "aaa"}))
	//fmt.Println(wordBreak3("aaaaaaa", []string{"aaaa", "aaa"}))
	//fmt.Println(wordBreak3("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak6("leetcode", []string{"leet", "code"}))
	//fmt.Println(wordBreak3("leetcode", []string{"leet", "code"}))
}

func Test_wordBreak7(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
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
				s:        "leetcode",
				wordDict: []string{"leet", "code"},
			},
			want: true,
		},
		{
			name: "t2",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak7(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak7() = %v, want %v", got, tt.want)
			}
		})
	}
}
