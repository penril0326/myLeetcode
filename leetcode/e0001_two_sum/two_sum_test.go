package twosum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want2 []int
	}{
		{
			name: "1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want:  []int{0, 1},
			want2: []int{1, 0},
		},
		{
			name: "2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want:  []int{1, 2},
			want2: []int{2, 1},
		},
		{
			name: "3",
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			want:  []int{0, 1},
			want2: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) && !reflect.DeepEqual(got, tt.want2) {
				t.Errorf("twoSum() = %v, want %v or want %v", got, tt.want, tt.want2)
			}
		})
	}
}
