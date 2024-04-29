package dynamicprogramming

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name:    "Case #1",
			args:    args{s: "babad"},
			wantAns: "aba",
		},
		{
			name:    "Case #2",
			args:    args{s: "cbbd"},
			wantAns: "bb",
		},
		{
			name:    "Case #3",
			args:    args{s: "a"},
			wantAns: "a",
		},
		{
			name:    "Case #4",
			args:    args{s: "ccc"},
			wantAns: "ccc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestPalindrome(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("longestPalindrome() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
