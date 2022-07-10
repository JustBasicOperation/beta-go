package algorithm

import (
	"testing"
)

func Test_twoSum(t *testing.T) {

}

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "case2",
			args: args{
				s: " ",
			},
			want: 1,
		},
		{
			name: "case3",
			args: args{
				s: "abcabab",
			},
			want: 3,
		},
		{
			name: "case3",
			args: args{
				s: "abc",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
