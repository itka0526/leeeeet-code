package bitmanipulation

import "testing"

func Test_hammingWeight(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Case #1",
			args:    args{n: 11},
			wantAns: 3,
		},
		{
			name:    "Case #2",
			args:    args{n: 128},
			wantAns: 1,
		},
		{
			name:    "Case #3",
			args:    args{n: 2147483645},
			wantAns: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := hammingWeight(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("hammingWeight() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
