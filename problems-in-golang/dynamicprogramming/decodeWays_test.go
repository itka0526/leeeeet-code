package dynamicprogramming

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Case #1",
			args:    args{s: "12"},
			wantAns: 2,
		},
		{
			name:    "Case #2",
			args:    args{s: "226"},
			wantAns: 3,
		},
		{
			name:    "Case #3",
			args:    args{s: "10"},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numDecodings(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("numDecodings() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
