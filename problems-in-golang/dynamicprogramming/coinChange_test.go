package dynamicprogramming

import "testing"

func TestCoinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case #1",
			args: args{coins: []int{1, 2, 5}, amount: 11},
			want: 3,
		},
		{
			name: "Case #2",
			args: args{coins: []int{1, 3, 4, 5}, amount: 7},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
