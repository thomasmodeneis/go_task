package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Given A = [3, 4, 5, 3, 7] your function should return 3",
			args: args{A: []int{3, 4, 5, 3, 7}},
			want: 3,
		},
		{
			name: "Given A = [1, 2, 3, 4] your function should return −1",
			args: args{A: []int{1, 2, 3, 4}},
			want: -1,
		},
		{
			name: "Given A = [1, 3, 1, 2] your function should return 0",
			args: args{A: []int{1, 3, 1, 2}},
			want: 0,
		},
		{
			name: "Given A = [2, 3, 1, 8, 4, 10, 8, 10, 6, 8] your function should return 0",
			args: args{A: []int{2, 3, 1, 8, 4, 10, 8, 10, 6, 8}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Solution(tt.args.A)
			require.Equal(t, tt.want, got)
		})
	}
}
