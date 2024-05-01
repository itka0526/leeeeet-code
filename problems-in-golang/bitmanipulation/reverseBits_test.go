package bitmanipulation

import "testing"

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name    string
		args    args
		wantAns uint32
	}{
		{
			name:    "Case #1",
			args:    args{num: 43261596},
			wantAns: 964176192,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := reverseBits(tt.args.num); gotAns != tt.wantAns {
				t.Errorf("reverseBits() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
