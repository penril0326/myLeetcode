package maxsumofapairwithequalsumofdigit

import "testing"

func Test_maximumSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{18, 43, 36, 13, 7},
			},
			want: 54,
		},
		{
			name: "2",
			args: args{
				nums: []int{18, 43, 36, 13, 7, 27},
			},
			want: 63,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSum(tt.args.nums); got != tt.want {
				t.Errorf("maximumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
